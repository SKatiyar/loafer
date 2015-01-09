package loafer

import (
	"time"
)

type level int
type Fields map[string]interface{}

const (
	Info level = iota
	Warn
	Error
	Fatal
)

type Loafer struct {
	UniqueId  string
	Adaptor   Adaptor
	Formatter Formatter
}

type Logger struct {
	id string
	a  Adaptor
	f  Formatter
	df Formatter
}

func (l *Logger) Log(e level, data Fields) {
	// get time at which log was generated
	ct := time.Now().UTC()

	// if time field exist then cascade it
	_, ok := data["time"]
	if ok {
		data["f.time"] = data["time"]
	}
	data["time"] = ct

	// if level field exist then cascade it
	_, ok = data["level"]
	if ok {
		data["f.level"] = data["level"]
	}
	data["level"] = e

	// if uid field exist then cascade it
	_, ok = data["uid"]
	if ok {
		data["f.uid"] = data["uid"]
	}
	data["uid"] = l.id

	// call the provided formatter, but if it fails resort to default
	formatedData, fErr := l.f.Format(data)
	if fErr != nil {
		formatedData, _ = l.df.Format(data)
	}

	_, _ = l.a.Write(formatedData)
}

func NewLoafer(l Loafer) *Logger {
	return &Logger{l.UniqueId, l.Adaptor, l.Formatter, &defaultFormatter{}}
}
