package yjt

import "testing"

func TestParseJson(t *testing.T) {
	if err := parseJson(); err != nil {
		t.Logf("Could not parse json %s", err)
		t.Fatal(err)
	}
}