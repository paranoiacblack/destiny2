package destiny2

import (
	"archive/zip"
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"golang.org/x/text/language"

	_ "github.com/mattn/go-sqlite3"
)

const apiPath = "https://www.bungie.net/Platform"

var supportedLanguages = language.NewMatcher([]language.Tag{
	language.English, // Used as a fallback if no language option is provided.
	language.Spanish,
	language.French,
	language.MustParse("es-mx"),
	language.German,
	language.Italian,
	language.Japanese,
	language.BrazilianPortuguese,
	language.Russian,
	language.Polish,
	language.Korean,
	language.TraditionalChinese,
	language.SimplifiedChinese,
})

// LocaleError represents an error finding a user-specified locale.
type LocaleError struct {
	locale string
}

func (e LocaleError) Error() string {
	return fmt.Sprintf("%q is unsupported by the Bungie API", e.locale)
}

// Manifest is a representation of DestinyManifest, the external-facing contract
// for just the properties needed by those calling the Destiny Platform API.
type Manifest struct {
	// Manifest version, used to check if an update is necessary.
	version string
	// Contracts are tables in the Manifest Database for specific Destiny2 entities.
	// For each language/locale, there is a different set of paths to the appropriate contract in the manifest.
	contracts map[language.Tag]map[string]string

	// Mobile contracts for each language/locale. Mobile contracts
	// are stored as SQL databases.
	mobileContracts       map[language.Tag]string
	cachedMobileManifests map[language.Tag]string

	// Content paths.
	cdn gearCDN

	// Lookup tables for information about gear and clan banners.
	gearAssetPath  []string
	clanBannerPath string
}

// NewManifest returns a populated Destiny 2 Manifest, similar to the
// Bungie API call Destiny2.GetDestinyManifest.
func NewManifest() (*Manifest, error) {
	m := new(Manifest)
	if _, err := m.Update(); err != nil {
		return nil, err
	}
	return m, nil
}

// Close closes this manifest, cleaning up any temporary files created during contract fulfillment.
func (m *Manifest) Close() {
	// For now, there's just potentially files representing mobile manifest DBs.
	m.version = ""
	m.contracts = nil
	m.mobileContracts = nil
	for _, db := range m.cachedMobileManifests {
		os.Remove(db)
	}
	m.cachedMobileManifests = nil
	m.cdn = gearCDN{}
	m.gearAssetPath = nil
	m.clanBannerPath = ""
}

// UpdateStatus is returned from Update to communicate whether or not the manifest was updated to a new version.
// If Successful, all fulfilled contracts are now outdated.
type UpdateStatus int

const (
	AlreadyUpdated UpdateStatus = iota
	Successful
	Failed
)

// Update updates the manifest to the newest version, if necessary.
func (m *Manifest) Update() (UpdateStatus, error) {
	resp, err := http.Get(apiPath + "/Destiny2/Manifest")
	if err != nil {
		return Failed, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Failed, err
	}

	// TODO(paranoiacblack): Handle error status and error code more traditionally.
	var manifestResp struct {
		Response        json.RawMessage
		ErrorCode       int
		ThrottleSeconds int
		ErrorStatus     string
		Message         string
		MessageData     json.RawMessage
	}

	if err := json.Unmarshal(body, &manifestResp); err != nil {
		return Failed, err
	}
	status, err := m.parseManifest(manifestResp.Response)
	if err != nil {
		return Failed, err
	}
	return status, nil
}

// Version returns the version string of this manifest.
func (m Manifest) Version() string {
	return m.version
}

// FulfillContract fulfills a Bungie.net contract by adding all related entities for a given definition.
func (m *Manifest) FulfillContract(definition Contract, opts ...FulfillmentOption) error {
	fulfillmentOpt := fulfillmentOptions{tag: language.English}
	for _, opt := range opts {
		if err := opt(&fulfillmentOpt); err != nil {
			return err
		}
	}

	retrievalFunc := m.retrieveFromComponent
	if fulfillmentOpt.mobile {
		retrievalFunc = m.retrieveFromMobile
	}

	data, err := retrievalFunc(fulfillmentOpt.tag, definition.Name())
	if err != nil {
		return err
	}
	return json.Unmarshal(data, definition)
}

func (m *Manifest) retrieveFromComponent(tag language.Tag, name string) ([]byte, error) {
	contractPath, ok := m.contracts[tag][name]
	if !ok {
		return nil, fmt.Errorf("%q is not a valid Destiny.Definitions name", name)
	}

	resp, err := http.Get(fmt.Sprintf("https://www.bungie.net%s", contractPath))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (m *Manifest) retrieveFromMobile(tag language.Tag, name string) ([]byte, error) {
	// Bungie's mobile manifest is stored at a given location for each language/locale
	// as a zipped sqlite DB. To save effort redownloading and unzipping the DB each time,
	// we'll write it to a temporary file and cache DB connections to it. So retrieving data
	// from mobile manifest should be fast on repeated contract fulfillments, but possibly slow
	// the first time.
	if len(m.cachedMobileManifests) == 0 {
		m.cachedMobileManifests = map[language.Tag]string{}
	}
	if _, ok := m.cachedMobileManifests[tag]; !ok {
		// First time getting the mobile manifest for this locale, or first time
		// since an update.
		if err := m.createTempDB(tag); err != nil {
			return nil, err
		}
	}
	dbPath := m.cachedMobileManifests[tag]
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// TODO(paranoiacblack): Not exactly sure, but name should probably be escaped here to avoid malicious queries.
	// At the same time, it's querying a temporary read-only database so it probably doesn't matter if a malicious
	// user tries anything.
	rows, err := db.QueryContext(context.Background(), fmt.Sprintf("SELECT * FROM %s", name))
	if err != nil {
		return nil, err
	}

	contract := map[uint32]json.RawMessage{}
	for rows.Next() {
		var key int
		var data []byte
		if err := rows.Scan(&key, &data); err != nil {
			return nil, err
		}
		contract[uint32(key)] = data
	}
	return json.Marshal(contract)
}

// createTempDB creates a temporary file with the sqlite destiny 2 mobile manifest.
func (m *Manifest) createTempDB(tag language.Tag) error {
	mobilePath, ok := m.mobileContracts[tag]
	if !ok {
		return fmt.Errorf("there is no existing mobile manifest for language/locale %s", tag)
	}

	resp, err := http.Get("https://www.bungie.net" + mobilePath)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	zippedDB := bytes.NewReader(body)
	r, err := zip.NewReader(zippedDB, zippedDB.Size())
	if err != nil {
		return err
	}

	f := r.File[0]
	fc, err := f.Open()
	if err != nil {
		return err
	}
	defer fc.Close()

	content, err := ioutil.ReadAll(fc)
	if err != nil {
		return err
	}

	tempDB, err := ioutil.TempFile("", "d2manifest")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(tempDB.Name(), content, 0644); err != nil {
		return err
	}
	m.cachedMobileManifests[tag] = tempDB.Name()
	return nil
}

type fulfillmentOptions struct {
	// tag is the supported tag for a given language/locale
	tag language.Tag
	// if true, use the mobile manifest when fulfilling a contract
	mobile bool
}

// FulfillmentOption is an optional way to fulfill a given contract.
type FulfillmentOption func(o *fulfillmentOptions) error

// WithLocale fulfills a contract using a specific language/locale
// supported by Bungie.
func WithLocale(locale string) FulfillmentOption {
	return func(o *fulfillmentOptions) error {
		o.tag = getSupportedTagForLocale(locale)
		if o.tag == language.Und {
			return LocaleError{locale}
		}
		return nil
	}
}

// UseMobileManifest fulfills a contract using data from the Mobile Manifest.
func UseMobileManifest(mobile bool) FulfillmentOption {
	return func(o *fulfillmentOptions) error {
		o.mobile = mobile
		return nil
	}
}

// Paths where specific rendering information can be found.
type gearCDN struct {
	Geometry, Texture, PlateRegion, Gear, Shader string
}

// The JSON structure of the manifest response data.
type manifestResponse struct {
	Version string
	// Omitted for now, the database appears to be empty.
	MobileAssetContentPath   string `json:"-"`
	MobileGearAssetDataBases []struct {
		Version int
		Path    string
	}
	MobileWorldContentPaths json.RawMessage

	// For the purpose of this API, we won't be using the aggregate
	// JSON files that describe all world content.
	JsonWorldContentPaths json.RawMessage `json:"-"`
	// Instead, we use the individual components which define each entity.
	JsonWorldComponentContentPaths json.RawMessage
	MobileClanBannerDatabasePath   string
	MobileGearCDN                  gearCDN

	// Unused as far as I can tell.
	IconImaginePyramidInfo []string `json:"-"`
}

func (m *Manifest) parseManifest(data []byte) (UpdateStatus, error) {
	var resp manifestResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return Failed, err
	}
	if m.version == resp.Version {
		// Manifest is already updated to latest version.
		return AlreadyUpdated, nil
	}

	// Close the current manifest, basically clear all the fields and cleanup.
	m.Close()

	if err := m.parseMobileContentPaths(resp.MobileWorldContentPaths); err != nil {
		return Failed, err
	}

	if err := m.parseContractPaths(resp.JsonWorldComponentContentPaths); err != nil {
		return Failed, err
	}

	m.version = resp.Version
	gearDBs := []string{}
	for _, db := range resp.MobileGearAssetDataBases {
		gearDBs = append(gearDBs, db.Path)
	}
	m.gearAssetPath = gearDBs
	m.clanBannerPath = resp.MobileClanBannerDatabasePath
	m.cdn = resp.MobileGearCDN

	return Successful, nil
}

// mobileWorldContentPaths are defined as {"locale": "path"}
func (m *Manifest) parseMobileContentPaths(data []byte) error {
	contentPaths := make(map[string]string)
	if err := json.Unmarshal(data, &contentPaths); err != nil {
		return err
	}

	m.mobileContracts = map[language.Tag]string{}
	for locale, path := range contentPaths {
		tag := getSupportedTagForLocale(locale)
		if tag == language.Und {
			return LocaleError{locale}
		}
		m.mobileContracts[tag] = path
	}
	return nil
}

// jsonWorldComponentsPath are defined as {"locale": {"definition": "path"}}
func (m *Manifest) parseContractPaths(data []byte) error {
	contractPaths := make(map[string]map[string]string)
	if err := json.Unmarshal(data, &contractPaths); err != nil {
		return err
	}

	m.contracts = map[language.Tag]map[string]string{}
	for locale, contract := range contractPaths {
		tag := getSupportedTagForLocale(locale)
		if tag == language.Und {
			return LocaleError{locale}
		}

		m.contracts[tag] = map[string]string{}
		for name, path := range contract {
			m.contracts[tag][name] = path
		}
	}
	return nil
}

// Returns a language.Tag that matches Bungie-supported locales.
func getSupportedTagForLocale(locale string) language.Tag {
	tag, _, confidence := supportedLanguages.Match(language.Make(locale))
	if confidence == language.No {
		switch locale {
		case "zh-cht":
			tag = language.TraditionalChinese
		case "zh-chs":
			tag = language.SimplifiedChinese
		}
	}
	return tag
}
