package yjt

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v2"
)

func YamlToJson(y []byte) ([]byte, error) {
	var yamlObj interface{}

	err := yaml.Unmarshal(y, &yamlObj)
	if err != nil {
		return nil, err
	}

	jsonObj, err := convertToJson(yamlObj)

	if err != nil {
		return nil, err
	}

	return json.Marshal(jsonObj)
}

func YamlToToml(y []byte) ([]byte, error) {
	var yamlObj interface{}
	err := yaml.Unmarshal(y, &yamlObj)
	if err != nil {
		return nil, err
	}

	jsonObj, err := convertToJson(yamlObj)

	if err != nil {
		return nil, err
	}

	return toml.Marshal(jsonObj)
}

func convertToJson(yObj interface{}) (interface{}, error) {
	var err error

	switch yamlType := yObj.(type) {
	case map[interface{}]interface{}:
		stringMap := make(map[string]interface{})
		for k, v := range yamlType {
			var keyString string
			switch keyType := k.(type) {
			case string:
				keyString = keyType
			case int:
				keyString = strconv.Itoa(keyType)
			case int64:
				keyString = strconv.FormatInt(keyType, 10)
			case float64:
				s := strconv.FormatFloat(keyType, 'g', -1, 32)
				switch s {
				case "+Inf":
					s = ".inf"
				case "-Inf":
					s = "-.inf"
				case "NaN":
					s = ".nan"
				}
				keyString = s
			case bool:
				if keyType {
					keyString = "true"
				} else {
					keyString = "false"
				}
			default:
				return nil, fmt.Errorf("unsupported map key of type: %s, key: %+#v, value: %+#v,", reflect.TypeOf(k), k, v)
			}
			stringMap[keyString], err = convertToJson(v)
			if err != nil {
				return nil, err
			}
		}
		return stringMap, nil
	case []interface{}:
		array := make([]interface{}, len(yamlType))
		for i, v := range yamlType {
			array[i], err = convertToJson(v)
			if err != nil {
				return nil, err
			}
		}
		return array, nil
	default:
		return yObj, nil
	}
}
