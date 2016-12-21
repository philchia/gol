package gol

import (
	"io"
)

// Logger ...
type Logger interface {
	Debug(i ...interface{})
	Debugf(format string, i ...interface{})

	Info(i ...interface{})
	Infof(format string, i ...interface{})

	Warn(i ...interface{})
	Warnf(format string, i ...interface{})

	Error(i ...interface{})
	Errorf(format string, i ...interface{})

	Critical(i ...interface{})
	Criticalf(format string, i ...interface{})
}

type gollog struct {
	golWriter io.Writer
}

func NewLogger() Logger {

}
