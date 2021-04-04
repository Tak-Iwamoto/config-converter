package yjt

import (
	_ "embed"
	"fmt"

	"gopkg.in/yaml.v2"
)

//go:embed test/test.yml
var testYml []byte

func parseYml() error {
	y := make(map[string]interface{})

	err := yaml.Unmarshal(testYml, &y)
	if err != nil {
		return err
	}

	fmt.Println(y)

	return nil
}
