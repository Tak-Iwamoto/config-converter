package yjt

import "testing"

func TestParseYml(t *testing.T) {
	if err := parseYml(); err != nil {
		t.Logf("Could not parse json %s", err)
		t.Fatal(err)
	}
}