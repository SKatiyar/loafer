package loafer

import (
	"testing"
)

func TestDefaultAdaptor(t *testing.T) {
	var data = []byte(`key:error`)

	a := &defaultAdaptor{}
	_, err := a.Write(data)
	if err != nil {
		t.Error(err)
	}
}
