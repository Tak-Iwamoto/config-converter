package yjt

import (
	"encoding/json"

	"github.com/pelletier/go-toml"
)

func TomlToJson(t []byte) ([]byte, error) {
	var jObj interface{}
	if err := toml.Unmarshal(t, &jObj); err != nil {
		return nil, err
	}

	return json.Marshal(jObj)
}

func JsonToToml(j []byte) ([]byte, error) {
	var tObj interface{}
	if err := json.Unmarshal(j, &tObj); err != nil {
		return nil, err
	}

	return toml.Marshal(tObj)
}
