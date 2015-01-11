package loafer

import (
	"encoding/json"
	"testing"
)

type testStruct struct {
	Key string `json:key`
}

var data Fields

func TestJsonFormatter(t *testing.T) {
	data = make(Fields)
	data["key"] = "error"

	f := &JSONFormatter{}
	j, err := f.Format(data)
	if err != nil {
		t.Error(err)
	}

	var s testStruct
	uErr := json.Unmarshal(j, &s)
	if uErr != nil {
		t.Error(uErr)
	}
	if s.Key != "error" {
		t.Error("Wrong formatting")
	}
}
