package yjt

import (
	"testing"
)

func TestYamlToJson(t *testing.T) {
	var tests = []struct {
		yaml string
		json string
	}{
		{"a: a\n",
			`{"a":"a"}`,
		},
		{"a: a\nb: b\n",
			`{"a":"a","b":"b"}`,
		},
		{"a: null\n",
			`{"a":null}`,
		},
		{"1000000000000000000000000000000000000: a\n",
			`{"1e+36":"a"}`,
		},
		{"1.2: a\n",
			`{"1.2":"a"}`,
		},
		{`[{a: a}, {a: {b: 1, c: 2}}]`,
			`[{"a":"a"},{"a":{"b":1,"c":2}}]`,
		},
		{"- t: \n",
			`[{"t":null}]`,
		},
	}

	for _, test := range tests {
		j, err := YamlToJson([]byte(test.yaml))
		if err != nil {
			t.Errorf("Failed to convert yaml to json, input %s, err: %v,", test.yaml, err.Error())
		}
		if test.json != string(j) {
			t.Errorf("Failed to convert yaml to json, expected `%s`, actual `%s`", test.json, string(j))
		}
	}
}

func TestYamlToToml(t *testing.T) {
	var tests = []struct {
		yaml string
		toml string
	}{
		{"t: t", "t = \"t\"\n"},
		{"t: true", "t = true\n"},
		{"t: 1.2", "t = 1.2\n"},
		{"person:\n  name: Tom\nt: t\n", "t = \"t\"\n\n[person]\n  name = \"Tom\"\n"},
		{"people:\n- age: 24\n  name: Tom\n- age: 27\n  name: Tak\n", "people = [{ age = 24, name = \"Tom\" }, { age = 27, name = \"Tak\" }]\n"},
	}
	for _, test := range tests {
		toml, err := YamlToToml([]byte(test.yaml))
		if err != nil {
			t.Errorf("Failed to convert yaml to toml detail: %s", err.Error())
		}

		if test.toml != string(toml) {
			t.Errorf("Failed to convert toml to yaml, expected: %s, actual: %s", test.toml, string(toml))
		}
	}
}
