package loafer

import (
	"fmt"
)

type FmtAdaptor struct{}

func (f *FmtAdaptor) Write(b []byte) (int, error) {
	return fmt.Println(string(b[:len(b)]))
	// Output:
	// key:error
}
