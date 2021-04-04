package yjt

import "testing"

type Case struct {
	input  string
	output string
}

func TestParseYml(t *testing.T) {
	cases := []Case{
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

	for _, c := range cases {
		result, err := YmlToJson([]byte(c.input))
		if err != nil {
			t.Errorf("Failed to convert %s, input %s, err: %v,", "YmlToJson", c.input, err)
		}
		if string(result) != c.output {
			t.Errorf("Failed to convert %s, input `%s`, expected `%s`, got `%s`", "YmlToJson", c.input, c.output, result)
		}
	}
}
