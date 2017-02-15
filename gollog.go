package gol

import (
	"fmt"

	"io"

	"github.com/philchia/gol/adapter"
	"github.com/philchia/gol/level"
)

// This works as a compiler check
var _ Logger = (*gollog)(nil)

type gollog struct {
	level    level.LogLevel
	option   LogOption
	adapters map[string]adapter.Adapter
	logChan  chan *logMSG
	doneChan chan struct{}
	// mutex    sync.RWMutex
}

func (l *gollog) msgPump() {

	for msg := range l.logChan {

		for _, v := range l.adapters {
			if v.Level() <= msg.logLevel {
				io.Copy(v, &msg.bf)
			}
		}
		msgPoolPut(msg)
	}

	close(l.doneChan)
}

func (l *gollog) put(msg *logMSG) {
	l.logChan <- msg
}

func (l *gollog) output(callDepth int, level level.LogLevel, msg string) {
	logmsg := msgPoolGet()
	logmsg.logLevel = level
	l.generateLog(callDepth, level, msg, &logmsg.bf)
	l.put(logmsg)
}

// Debug will prinnt log as DEBUG level
func (l *gollog) Debug(i ...interface{}) {
	if l.level > level.DEBUG {
		return
	}
	l.output(2, level.DEBUG, fmt.Sprint(i...))
}

// Debugf will prinnt log as DEBUG level
func (l *gollog) Debugf(format string, i ...interface{}) {
	if l.level > level.DEBUG {
		return
	}
	l.output(2, level.DEBUG, fmt.Sprintf(format, i...))
}

// Info will prinnt log as INFO level
func (l *gollog) Info(i ...interface{}) {
	if l.level > level.INFO {
		return
	}
	l.output(2, level.INFO, fmt.Sprint(i...))
}

// Infof will prinnt log as INFO level
func (l *gollog) Infof(format string, i ...interface{}) {
	if l.level > level.INFO {
		return
	}
	l.output(2, level.INFO, fmt.Sprintf(format, i...))
}

// Warn will prinnt log as WARN level
func (l *gollog) Warn(i ...interface{}) {
	if l.level > level.WARN {
		return
	}
	l.output(2, level.WARN, fmt.Sprint(i...))
}

// Warnf will prinnt log as WARN level
func (l *gollog) Warnf(format string, i ...interface{}) {
	if l.level > level.WARN {
		return
	}
	l.output(2, level.WARN, fmt.Sprintf(format, i...))
}

// Error will prinnt log as ERROR level
func (l *gollog) Error(i ...interface{}) {
	if l.level > level.ERROR {
		return
	}
	l.output(2, level.ERROR, fmt.Sprint(i...))
}

// Errorf will prinnt log as ERROR level
func (l *gollog) Errorf(format string, i ...interface{}) {
	if l.level > level.ERROR {
		return
	}
	l.output(2, level.ERROR, fmt.Sprintf(format, i...))
}

// Critical will prinnt log as CRITICAL level
func (l *gollog) Critical(i ...interface{}) {
	if l.level > level.CRITICAL {
		return
	}
	l.output(2, level.CRITICAL, fmt.Sprint(i...))
}

// Criticalf will prinnt log as CRITICAL level
func (l *gollog) Criticalf(format string, i ...interface{}) {
	if l.level > level.CRITICAL {
		return
	}
	l.output(2, level.CRITICAL, fmt.Sprintf(format, i...))
}

// Flush flush all buffered log and call Close() on all adapters
func (l *gollog) Flush() {
	close(l.logChan)
	<-l.doneChan
	for _, c := range l.adapters {
		c.Close()
	}
}
