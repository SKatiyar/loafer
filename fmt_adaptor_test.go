package loafer

import (
	"testing"
)

func TestFmtAdaptor(t *testing.T) {
	var data = []byte(`key:error`)

	a := &FmtAdaptor{}
	_, err := a.Write(data)
	if err != nil {
		t.Error(err)
	}
}
