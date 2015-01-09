package loafer

import (
	"fmt"
)

type Formatter interface {
	Format(Fields) ([]byte, error)
}

type defaultFormatter struct{}

// TODO fix the colour rendering
var colorArr []string = []string{"\x1b[32;1m ", "\x1b[33;1m ", "\x1b[35;1m ", "\x1b[31;1m "}

func (d *defaultFormatter) Format(f Fields) ([]byte, error) {
	var str string
	if val, ok := f["level"]; ok {
		str += colorArr[val.(level)]
	}
	for key, value := range f {
		str += fmt.Sprintf("[ %s ] %+v ", key, value)
	}

	str += "\x1b[0m"

	return []byte(str), nil
}
