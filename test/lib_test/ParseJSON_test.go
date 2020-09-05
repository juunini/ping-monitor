package lib_test

import (
	"ping-monitor/lib"
	"testing"
)

func TestParseJSON(t *testing.T) {
	var decoder struct {
		Test     bool   `json:"test"`
		Language string `json:"language"`
	}

	err := lib.ParseJSON("./ParseJSON_test.json", &decoder)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(decoder.Language, decoder.Test)
}
