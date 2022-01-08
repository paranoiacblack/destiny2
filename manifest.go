package destiny2

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/text/language"
)

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

	// Mobile contracts for each language/locale.
	mobileContracts map[language.Tag]string

	// Content paths.
	cdn gearCDN

	// Lookup tables for information about gear and clan banners.
	gearAssetPath  []string
	clanBannerPath string

	contractReader ContractReader
}

// NewManifest returns a populated Destiny 2 Manifest, similar to the
func NewManifest(reader ContractReader) (*Manifest, error) {
	m := &Manifest{contractReader: reader}
	if err := m.Update(nil); err != nil {
		return nil, err
	}
	return m, nil
}

// UpdateFunc is a closure that is run after a successful update.
type UpdateFunc func() error

// Update updates the manifest to the newest version and, if necessary, runs updateFn.
func (m *Manifest) Update(updateFn UpdateFunc) error {
	resp, err := http.Get("https://www.bungie.net/Platform/Destiny2/Manifest")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
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
		return err
	}
	if err := m.parseManifest(manifestResp.Response, updateFn); err != nil {
		return err
	}
	return nil
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

	tag := fulfillmentOpt.tag
	path, ok := m.contracts[tag][definition.Name()]
	if fulfillmentOpt.mobile {
		path, ok = m.mobileContracts[tag]
	}
	if !ok {
		return fmt.Errorf("%q is not a valid Destiny.Definitions name", definition.Name())
	}

	data, err := m.contractReader.ReadContract(definition, path, fulfillmentOpt.mobile)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, definition)
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

func (m *Manifest) parseManifest(data []byte, updateFn UpdateFunc) error {
	var resp manifestResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		return err
	}
	if m.version == resp.Version {
		// Manifest is already updated to latest version.
		return nil
	}

	if err := m.parseMobileContentPaths(resp.MobileWorldContentPaths); err != nil {
		return err
	}

	if err := m.parseContractPaths(resp.JsonWorldComponentContentPaths); err != nil {
		return err
	}

	m.version = resp.Version
	gearDBs := []string{}
	for _, db := range resp.MobileGearAssetDataBases {
		gearDBs = append(gearDBs, db.Path)
	}
	m.gearAssetPath = gearDBs
	m.clanBannerPath = resp.MobileClanBannerDatabasePath
	m.cdn = resp.MobileGearCDN

	if updateFn != nil {
		if err := updateFn(); err != nil {
			return err
		}
	}
	return nil
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
