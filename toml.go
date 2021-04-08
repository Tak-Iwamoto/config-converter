package yjt

import (
	"encoding/json"

	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v2"
)

func TomlToJson(t []byte) ([]byte, error) {
	var jsonObj interface{}
	if err := toml.Unmarshal(t, &jsonObj); err != nil {
		return nil, err
	}

	return json.Marshal(jsonObj)
}

func TomlToYaml(t []byte) ([]byte, error) {
	var yamlObj interface{}
	if err := toml.Unmarshal(t, &yamlObj); err != nil {
		return nil, err
	}

	return yaml.Marshal(yamlObj)
}
