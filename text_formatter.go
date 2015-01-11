package loafer

import (
	"fmt"
)

type TextFormatter struct{}

// TODO fix the colour rendering
var colorArr []string = []string{"\x1b[32;1m ", "\x1b[33;1m ", "\x1b[35;1m ", "\x1b[31;1m "}

func (d *TextFormatter) Format(f Fields) ([]byte, error) {
	var str string
	if val, ok := f["level"]; ok {
		str += colorArr[val.(level)]
	}

	if tUid, ok := f["uid"]; ok {
		str += fmt.Sprintf("[ %s ] %+v ", "uid", tUid)
		delete(f, "uid")
	}

	if tLvl, ok := f["level"]; ok {
		str += fmt.Sprintf("[ %s ] %+v ", "level", tLvl)
		delete(f, "level")
	}

	if tVal, ok := f["time"]; ok {
		str += fmt.Sprintf("[ %s ] %+v ", "time", tVal)
		delete(f, "time")
	}

	for key, value := range f {
		str += fmt.Sprintf("[ %s ] %+v ", key, value)
	}

	str += "\x1b[0m"

	return []byte(str), nil
}
