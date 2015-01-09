package loafer

import (
	"bytes"
	"fmt"
	"testing"
)

type testAdaptor struct{}

func (t *testAdaptor) Write(p []byte) (n int, err error) {
	num := bytes.NewReader(p).Len()
	s := string(p[:num])
	fmt.Println(s)
	// Output:
	// key:error

	return num, nil
}

type testFomatter struct{}

func (t *testFomatter) Format(data Fields) ([]byte, error) {
	b := &bytes.Buffer{}
	fmt.Fprintf(b, "%s:%s", "key", "error")
	b.WriteByte('\n')
	return b.Bytes(), nil
}

func TestLoafer(t *testing.T) {
	log := NewLoafer(Loafer{"1", &testAdaptor{}, &testFomatter{}})
	log.Log(Info, Fields{
		"key": "error",
	})
}
