package yjt

import (
	_ "embed"
	"encoding/json"
)

//go:embed test/test.json
var testJson []byte

func parseJson() error {
	j := make(map[string]interface{})

	err := json.Unmarshal(testJson, &j)
	if err != nil {
		return err
	}

	return nil
}
