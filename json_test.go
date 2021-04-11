package yjt

import (
	"testing"
)

func TestJsonToYaml(t *testing.T) {
	var tests = []struct {
		json string
		yaml string
	}{
		{
			`{"a":"a"}`, "a: a\n",
		},
		{
			`{"a":"a","b":"b"}`,
			"a: a\nb: b\n",
		},
		{
			`{"a":null}`,
			"a: null\n",
		},
		{
			`{"1e+36":"a"}`,
			"\"1e+36\": a\n",
		},
		{
			`{"1.2":"a"}`,
			"\"1.2\": a\n",
		},
		{
			`{"people":[{"age":24,"name":"Tom"},{"age":27,"name":"Tak"}]}`,
			"people:\n- age: 24\n  name: Tom\n- age: 27\n  name: Tak\n",
		},
		{
			`[{"t":null}]`,
			"- t: null\n",
		},
	}

	for _, test := range tests {
		y, err := JsonToYaml([]byte(test.json))
		if err != nil {
			t.Errorf("Failed to convert json to yaml detail: %s", err)
		}

		if test.yaml != string(y) {
			t.Errorf("Failed to convert json to yaml detail: %s, expected: %s, actual: %s", err, test.yaml, string(y))
		}
	}
}

func TestJsonToToml(t *testing.T) {
	var tests = []struct {
		json string
		toml string
	}{
		{`{"t":"t"}`, "t = \"t\"\n"},
		{`{"t":true}`, "t = true\n"},
		{`{"t":1.2}`, "t = 1.2\n"},
		{`{"person":{"name":"Tom"},"t":"t"}`, "t = \"t\"\n\n[person]\n  name = \"Tom\"\n"},
		{`{"people":[{"age":24,"name":"Tom"},{"age":27,"name":"Tak"}]}`, "people = [{ age = 24.0, name = \"Tom\" }, { age = 27.0, name = \"Tak\" }]\n"},
	}

	for _, test := range tests {
		toml, err := JsonToToml([]byte(test.json))
		if err != nil {
			t.Errorf("Failed to convert json to toml, detail: %s", err.Error())
		}

		if test.toml != string(toml) {
			t.Errorf("Failed to convert json to toml, expected: %s, actual: %s", test.toml, string(toml))
		}
	}
}
