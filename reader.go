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

	_ "github.com/mattn/go-sqlite3"
)

// ContractReader reads a Bungie.net contract.
type ContractReader interface {
	// ReadContract returns the marshalled contract with a given path.
	ReadContract(contract Contract, path string, useMobile bool) ([]byte, error)
	// Close closes this reader and all associated resources.
	Close() error
}

// BungieAPIReader reads contracts from the Bungie.net API.
type BungieAPIReader struct {
	cachedMobileManifests map[string]string
}

// ReadContract reads a contract at path from the Bungie.net endpoint.
func (r *BungieAPIReader) ReadContract(contract Contract, path string, useMobile bool) ([]byte, error) {
	if useMobile {
		return r.fromMobile(contract, path)
	}

	resp, err := http.Get(fmt.Sprintf("https://www.bungie.net%s", path))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func (r *BungieAPIReader) fromMobile(contract Contract, path string) ([]byte, error) {
	// Bungie's mobile manifest is stored at a given location for each language/locale
	// as a zipped sqlite DB. To save effort redownloading and unzipping the DB each time,
	// we'll write it to a temporary file and cache DB connections to it. So retrieving data
	// from mobile manifest should be fast on repeated contract fulfillments, but possibly slow
	// the first time.
	if len(r.cachedMobileManifests) == 0 {
		r.cachedMobileManifests = map[string]string{}
	}

	cachedName := contract.Name() + path
	if _, ok := r.cachedMobileManifests[cachedName]; !ok {
		// First time getting the mobile manifest for this contract path.
		if err := r.createTempDB(cachedName, path); err != nil {
			return nil, err
		}
	}

	dbPath := r.cachedMobileManifests[cachedName]
	data, err := readContractFromDB(dbPath, contract.Name())
	if err != nil {
		return nil, err
	}
	return data, nil
}

// createTempDB creates a temporary file with the sqlite destiny 2 mobile manifest.
func (r *BungieAPIReader) createTempDB(name, path string) error {
	resp, err := http.Get("https://www.bungie.net" + path)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	zippedDB := bytes.NewReader(body)
	zipReader, err := zip.NewReader(zippedDB, zippedDB.Size())
	if err != nil {
		return err
	}

	f := zipReader.File[0]
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
	r.cachedMobileManifests[name] = tempDB.Name()
	return nil
}

// Close removes temporary mobile databases.
func (r *BungieAPIReader) Close() error {
	for _, db := range r.cachedMobileManifests {
		if err := os.Remove(db); err != nil {
			return err
		}
	}
	return nil
}

// readContractFromDB is a helper for reading a sqlite DB containing Bungie.net contract.
func readContractFromDB(path string, contractName string) ([]byte, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.QueryContext(context.Background(), fmt.Sprintf("SELECT * FROM %s", contractName))
	if err != nil {
		return nil, err
	}

	definitions := map[uint32]json.RawMessage{}
	for rows.Next() {
		var key int
		var data []byte
		if err := rows.Scan(&key, &data); err != nil {
			return nil, err
		}
		definitions[uint32(key)] = data
	}
	return json.Marshal(definitions)
}
