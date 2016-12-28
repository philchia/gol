package gol

import (
	"fmt"

	"bytes"

	"io"

	"github.com/philchia/gol/adapter"
)

// This works as a compiler check
var _ Logger = (*gollog)(nil)

type gollog struct {
	level    LogLevel
	option   LogOption
	adapters map[string]adapter.Adapter
	logChan  chan *bytes.Buffer
	doneChan chan struct{}
	// mutex    sync.RWMutex
}

func (l *gollog) msgPump() {

	for buf := range l.logChan {
		for _, v := range l.adapters {
			io.Copy(v, buf)
		}
		bufferPoolPut(buf)
	}

	close(l.doneChan)
}

func (l *gollog) put(buf *bytes.Buffer) {
	l.logChan <- buf
}

func (l *gollog) output(callDepth int, level LogLevel, msg string) {
	buf := bufferPoolGet()
	l.generateLog(buf, callDepth, level, msg)

	l.put(buf)
}

// Debug will prinnt log as DEBUG level
func (l *gollog) Debug(i ...interface{}) {
	if l.level > DEBUG {
		return
	}
	l.output(2, DEBUG, fmt.Sprint(i...))
}

// Debugf will prinnt log as DEBUG level
func (l *gollog) Debugf(format string, i ...interface{}) {
	if l.level > DEBUG {
		return
	}
	l.output(2, DEBUG, fmt.Sprintf(format, i...))
}

// Info will prinnt log as INFO level
func (l *gollog) Info(i ...interface{}) {
	if l.level > INFO {
		return
	}
	l.output(2, INFO, fmt.Sprint(i...))
}

// Infof will prinnt log as INFO level
func (l *gollog) Infof(format string, i ...interface{}) {
	if l.level > INFO {
		return
	}
	l.output(2, INFO, fmt.Sprintf(format, i...))
}

// Warn will prinnt log as WARN level
func (l *gollog) Warn(i ...interface{}) {
	if l.level > WARN {
		return
	}
	l.output(2, WARN, fmt.Sprint(i...))
}

// Warnf will prinnt log as WARN level
func (l *gollog) Warnf(format string, i ...interface{}) {
	if l.level > WARN {
		return
	}
	l.output(2, WARN, fmt.Sprintf(format, i...))
}

// Error will prinnt log as ERROR level
func (l *gollog) Error(i ...interface{}) {
	if l.level > ERROR {
		return
	}
	l.output(2, ERROR, fmt.Sprint(i...))
}

// Errorf will prinnt log as ERROR level
func (l *gollog) Errorf(format string, i ...interface{}) {
	if l.level > ERROR {
		return
	}
	l.output(2, ERROR, fmt.Sprintf(format, i...))
}

// Critical will prinnt log as CRITICAL level
func (l *gollog) Critical(i ...interface{}) {
	if l.level > CRITICAL {
		return
	}
	l.output(2, CRITICAL, fmt.Sprint(i...))
}

// Criticalf will prinnt log as CRITICAL level
func (l *gollog) Criticalf(format string, i ...interface{}) {
	if l.level > CRITICAL {
		return
	}
	l.output(2, CRITICAL, fmt.Sprintf(format, i...))
}

// Flush flush all buffered log and call Close() on all adapters
func (l *gollog) Flush() {
	close(l.logChan)
	<-l.doneChan
	for _, c := range l.adapters {
		c.Close()
	}
}
