package yjt

import (
	"encoding/json"

	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v2"
)

func JsonToToml(j []byte) ([]byte, error) {
	var tomlObj interface{}
	if err := json.Unmarshal(j, &tomlObj); err != nil {
		return nil, err
	}

	return toml.Marshal(tomlObj)
}

func JsonToYaml(j []byte) ([]byte, error) {
	var jsonObj interface{}

	err := json.Unmarshal(j, &jsonObj)
	if err != nil {
		return nil, err
	}
	return yaml.Marshal(jsonObj)
}
