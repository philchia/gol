package gol

import (
	"os"

	"os/signal"

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
	AddLogAdapter(adapter.Adapter)
}

// NewLogger create a Logger
func NewLogger(level LogLevel) Logger {
	logger := &gollog{
		level:      level,
		option:     LstdFlags,
		logChan:    make(chan string, 1024),
		doneChan:   make(chan struct{}),
		signalChan: make(chan os.Signal, 1),
	}

	signal.Notify(logger.signalChan, os.Interrupt, os.Kill)
	logger.AddLogAdapter(console.NewAdapter())
	go func() {
		select {
		case <-logger.signalChan:
			logger.flush()
		}
	}()

	go logger.msgPump()
	return logger
}
