package loafer

import (
	"bytes"
	"testing"
	"time"
)

type temp struct {
	a int
}

func TestDefaultFormatter(t *testing.T) {
	f := &defaultFormatter{}
	data := Fields{
		"level":  level(3),
		"key":    "error",
		"you":    1,
		"time":   time.Now().UTC(),
		"struct": temp{1},
	}
	sb, _ := f.Format(data)
	str := bytes.NewBuffer(sb).String()
	if len(str) == 0 {
		t.Error("Not Formatted properly")
	}
}
