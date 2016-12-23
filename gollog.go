package gol

import (
	"fmt"
	"runtime"
	"time"

	"bytes"

	"errors"

	"sync"

	"github.com/philchia/gol/adapter"
)

// This works as a compiler check
var _ Logger = (*gollog)(nil)

type gollog struct {
	level    LogLevel
	option   LogOption
	adapters map[string]adapter.Adapter
	logChan  chan []byte
	doneChan chan struct{}
	mutex    sync.RWMutex
}

func (l *gollog) msgPump() {

	for msg := range l.logChan {
		l.mutex.RLock()
		for k := range l.adapters {
			l.adapters[k].Write(msg)
		}
		l.mutex.RUnlock()
	}

	close(l.doneChan)
}

func (l *gollog) put(msg []byte) {
	l.logChan <- msg
}

// itoa: Cheap integer to fixed-width decimal ASCII.  Give a negative width to avoid zero-padding.
func itoa(buf *bytes.Buffer, i int, wid int) {
	// Assemble decimal in reverse order.
	var b [20]byte
	bp := len(b) - 1
	for i >= 10 || wid > 1 {
		wid--
		q := i / 10
		b[bp] = byte('0' + i - q*10)
		bp--
		i = q
	}
	// i < 10
	b[bp] = byte('0' + i)
	buf.Write(b[bp:])
}

func (l *gollog) generatePrefix(buf *bytes.Buffer, callDepth int) {

	if l.option&(Ldate|Ltime|Lmicroseconds) != 0 {
		var t = time.Now()
		if l.option&LUTC != 0 {
			t = t.UTC()
		}
		if l.option&Ldate != 0 {
			year, month, day := t.Date()
			itoa(buf, year, 4)
			buf.WriteByte('/')
			itoa(buf, int(month), 2)
			buf.WriteByte('/')
			itoa(buf, day, 2)
			buf.WriteByte(' ')
		}
		if l.option&(Ltime|Lmicroseconds) != 0 {
			hour, min, sec := t.Clock()
			itoa(buf, hour, 2)
			buf.WriteByte(':')
			itoa(buf, min, 2)
			buf.WriteByte(':')
			itoa(buf, sec, 2)
			if l.option&Lmicroseconds != 0 {
				buf.WriteByte('.')
				itoa(buf, t.Nanosecond()/1e3, 6)
			}
			buf.WriteByte(' ')
		}
	}

	if l.option&(Lshortfile|Llongfile) != 0 {
		var file string
		var line int
		var ok bool
		_, file, line, ok = runtime.Caller(callDepth)
		if !ok {
			file = "???"
			line = 0
		}
		if l.option&Lshortfile != 0 {
			short := file
			for i := len(file) - 1; i > 0; i-- {
				if file[i] == '/' {
					short = file[i+1:]
					break
				}
			}
			file = short
		}
		buf.WriteString(file)
		buf.WriteByte(':')
		itoa(buf, line, -1)
		buf.WriteString(": ")
	}

}

func (l *gollog) generateLog(buf *bytes.Buffer, callDepth int, level LogLevel, msg string) {
	l.generatePrefix(buf, callDepth)

	buf.WriteString(level.ColorString())
	buf.WriteString(level.String())
	buf.WriteString(ALL.ColorString())
	buf.WriteByte(' ')
	buf.WriteString(msg)
	buf.WriteByte('\n')
}

func (l *gollog) output(callDepth int, level LogLevel, msg string) {
	buf := bufferPoolGet()
	buf.Reset()
	l.generateLog(buf, callDepth, level, msg)
	bts := buf.Bytes()
	bufferPoolPut(buf)
	l.put(bts)
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

// SetLevel set the shared logger's log level
func (l *gollog) SetLevel(level LogLevel) {
	l.level = level
}

// SetOption set the shared logger's log options used to format log headerr
func (l *gollog) SetOption(option LogOption) {
	l.option = option
}

// AddLogAdapter add a log adapter which implement the adapter.Adapter interface with give name key, return error if name already exists
func (l *gollog) AddLogAdapter(name string, adapter adapter.Adapter) error {
	l.mutex.Lock()
	if _, ok := l.adapters[name]; ok {
		l.mutex.Unlock()
		return errors.New("Adapter already exists")
	}
	l.adapters[name] = adapter
	l.mutex.Unlock()
	return nil
}

// RemoveAdapter remove a log adapter with give name key, return error in name not exists
func (l *gollog) RemoveAdapter(name string) error {
	l.mutex.Lock()
	if _, ok := l.adapters[name]; !ok {
		l.mutex.Unlock()
		return errors.New("Adapter not exists")
	}
	delete(l.adapters, name)
	l.mutex.Unlock()
	return nil
}

// Flush flush all buffered log and call Close() on all adapters
func (l *gollog) Flush() {
	close(l.logChan)
	<-l.doneChan
	for _, c := range l.adapters {
		c.Close()
	}
}
