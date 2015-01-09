package loafer

import (
	"fmt"
)

type Adaptor interface {
	Write([]byte) (int, error)
}

type defaultAdaptor struct{}

func (d *defaultAdaptor) Write(b []byte) (int, error) {
	return fmt.Println(string(b[:len(b)]))
	// Output:
	// key:error
}
