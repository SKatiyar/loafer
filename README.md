## Loafer
A small comfortable logger lib for golang

```go
package main

import (
    // import loafer to start logging
	"github.com/SKatiyar/loafer"
)

func main() {
    // create a new logging instance
	log := loafer.NewLoafer(loafer.Loafer{
		"1",                     // attach uniqueId to identify loafer instance
		&loafer.FmtAdaptor{},    // attach adaptor for sending logs
		&loafer.JSONFormatter{}, // attach formatter to format data
	})
	
	log.Log(loafer.Error, loafer.Fields{
		"loafer": "let's log this shit",
	})
}

```
