package yjt

import "testing"

func TestTomlToJson(t *testing.T) {
	var tests = []struct {
		toml string
		json string
	}{
		{"t = 't'", `{"t":"t"}`},
		{"t = true", `{"t":true}`},
		{"t = 1.2", `{"t":1.2}`},
		{"t = 't'\n[person]\nname='Tom'", `{"person":{"name":"Tom"},"t":"t"}`},
		{"[[people]]\nname='Tom'\nage=24\n[[people]]\nname='Tak'\nage=27", `{"people":[{"age":24,"name":"Tom"},{"age":27,"name":"Tak"}]}`},
	}

	for _, test := range tests {
		j, err := TomlToJson([]byte(test.toml))
		if err != nil {
			t.Errorf("Failed to convert toml to json detail: %s", err)
		}

		if test.json != string(j) {
			t.Errorf("Failed to convert toml to json detail: %s, expected: %s, actual: %s", err, test.json, string(j))
		}
	}
}
