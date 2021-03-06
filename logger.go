package logger

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

// FieldLogger interface
type FieldLogger interface {
	Logger
	WithField(string, interface{}) FieldLogger
	WithFields(map[string]interface{}) FieldLogger
}


type Logger interface {
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Printf(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
	Fatalf(string, ...interface{})
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})
	Fatal(...interface{})
	Panic(...interface{})
}


func ParseLevel(level string) (Level, error) {
	l, err := logrus.ParseLevel(level)
	return Level(l), err
}

// NewLogger based on the specified log level, defaults to "debug".
// See `New` for more details.
func NewLogger(level string) FieldLogger {
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		lvl = logrus.DebugLevel
	}
	return New(lvl)
}

// New based on the specified log level, defaults to "debug".
// This logger will log to the STDOUT in a human readable,
// but parseable form.
/*
	Example: time="2016-12-01T21:02:07-05:00" level=info duration=225.283µs human_size="106 B" method=GET path="/" render=199.79µs request_id=2265736089 size=106 status=200
*/
func New(lvl Level) FieldLogger {
	e := os.Getenv("GO_ENV")
	if len(e) == 0 {
		e = "development"
	}else if lvl == InfoLevel {
		e = "production"
	}
	dev := e == "development"

	l := logrus.New()
	l.SetOutput(os.Stdout)
	l.Level = lvl
	l.Formatter = &textFormatter{
		ForceColors: dev,
	}
	return Logrus{l}
}


func (l *StructuredLog) OnDebug(logging func()){
	if l.level == DebugLevel {
		logging()
	}
}
func JSONMarshal(structValue interface{}) string {
	s, _ := json.Marshal(structValue)
	return string(s)
}

func Struct(structValue interface{}) string{
	return fmt.Sprintf("%+v", structValue)
}