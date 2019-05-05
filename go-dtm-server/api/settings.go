package api

import (
	"encoding/json"
	"io/ioutil"
)

const _settingsPath = "./data/settings.json"

var _instance *Settings

// settings of project (singleton)
type Settings struct {
	ConnectionString string `json: "connectionString"`
	Connector        string `json: "connector"`
	Port             string `json: "port"`
}

// accesor for settigs instance
func GetSettings() (*Settings, error) {
	if _instance != nil {
		return _instance, nil
	}

	settings := &Settings{}

	bytes, err := ioutil.ReadFile(_settingsPath)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, settings)
	if err != nil {
		return nil, err
	}

	return settings, nil
}

func (this *Settings) Save() error {
	bytes, err := json.Marshal(this)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(_settingsPath, bytes, 0777)
	return err
}
