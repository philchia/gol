package gol

import (
	"bytes"

	"github.com/philchia/gol/adapter"
	"github.com/philchia/gol/adapter/console"
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

	SetLevel(LogLevel)
	SetOption(LogOption)

	AddLogAdapter(name string, adapter adapter.Adapter) error
	RemoveAdapter(name string) error
	Flush()
}

// CONSOLELOGGER represent the given console adapter name
const CONSOLELOGGER = "console"

// NewLogger create a Logger with given log level
func NewLogger(level LogLevel) Logger {
	logger := &gollog{
		level:    level,
		option:   LstdFlags,
		logChan:  make(chan *bytes.Buffer, 10240),
		doneChan: make(chan struct{}),
		adapters: make(map[string]adapter.Adapter, 1),
	}

	logger.AddLogAdapter(CONSOLELOGGER, console.NewAdapter())

	go logger.msgPump()
	return logger
}
