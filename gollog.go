package gol

import (
	"fmt"

	"github.com/philchia/gol/adapter"
	"github.com/philchia/gol/internal/stringUtil"
)

type gollog struct {
	level    LogLevel
	option   LogOption
	adapters []adapter.Adapter
	logChan  chan string
}

func (l *gollog) msgPump() {
	for {
		select {
		case msg := <-l.logChan:
			for _, adap := range l.adapters {
				go adap.Write([]byte(msg))
			}
		}
	}
}

func (l *gollog) put(msg string) {
	l.logChan <- msg
}

func (l *gollog) generateLog(level LogLevel, msg string) string {
	return stringUtil.JoinStrings("[", level.String(), "] ", msg, "\n")
}

func (l *gollog) Debug(i ...interface{}) {
	if l.level > DEBUG {
		return
	}
	msg := l.generateLog(DEBUG, fmt.Sprint(i...))
	l.put(msg)
}

func (l *gollog) Debugf(format string, i ...interface{}) {
	if l.level > DEBUG {
		return
	}
	msg := l.generateLog(DEBUG, fmt.Sprintf(format, i...))
	l.put(msg)
}

func (l *gollog) Info(i ...interface{}) {
	if l.level > INFO {
		return
	}
	msg := l.generateLog(INFO, fmt.Sprint(i...))
	l.put(msg)
}

func (l *gollog) Infof(format string, i ...interface{}) {
	if l.level > INFO {
		return
	}
	msg := l.generateLog(INFO, fmt.Sprintf(format, i...))
	l.put(msg)
}

func (l *gollog) Warn(i ...interface{}) {
	if l.level > WARN {
		return
	}
	msg := l.generateLog(WARN, fmt.Sprint(i...))
	l.put(msg)
}

func (l *gollog) Warnf(format string, i ...interface{}) {
	if l.level > WARN {
		return
	}
	msg := l.generateLog(WARN, fmt.Sprintf(format, i...))
	l.put(msg)
}

func (l *gollog) Error(i ...interface{}) {
	if l.level > ERROR {
		return
	}
	msg := l.generateLog(ERROR, fmt.Sprint(i...))
	l.put(msg)
}

func (l *gollog) Errorf(format string, i ...interface{}) {
	if l.level > ERROR {
		return
	}
	msg := l.generateLog(ERROR, fmt.Sprintf(format, i...))
	l.put(msg)
}

func (l *gollog) Critical(i ...interface{}) {
	if l.level > CRITICAL {
		return
	}
	msg := l.generateLog(CRITICAL, fmt.Sprint(i...))
	l.put(msg)
}

func (l *gollog) Criticalf(format string, i ...interface{}) {
	if l.level > CRITICAL {
		return
	}
	msg := l.generateLog(CRITICAL, fmt.Sprintf(format, i...))
	l.put(msg)
}

func (l *gollog) SetLevel(level LogLevel) {
	l.level = level
}

func (l *gollog) SetOption(option LogOption) {
	l.option = option
}

func (l *gollog) AddLogAdapter(a adapter.Adapter) {
	if a != nil {
		l.adapters = append(l.adapters, a)
	}
}
