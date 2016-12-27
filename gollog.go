package gol

import (
	"fmt"
	"runtime"
	"time"

	"bytes"

	"errors"

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

// itoa: Cheap integer to fixed-width decimal ASCII.  Give a negative width to avoid zero-padding.
func itoa(buf *bytes.Buffer, i int, wid int) {
	// Assemble decimal in reverse order.
	var b [8]byte
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

	buf.Write(level.ColorString())
	buf.Write(level.String())
	buf.Write(ALL.ColorString())
	buf.WriteByte(' ')
	buf.WriteString(msg)
	buf.WriteByte('\n')
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

// SetLevel set the shared logger's log level
func (l *gollog) SetLevel(level LogLevel) {
	l.level = level
}

// SetOption set the shared logger's log options used to format log headerr
func (l *gollog) SetOption(option LogOption) {
	l.option = option
}

// AddLogAdapter add a log adapter which implement the adapter.Adapter interface with give name key, return error if name already exists
func (l *gollog) AddLogAdapter(name string, adp adapter.Adapter) error {
	if adp == nil {
		return errors.New("nil adapter")
	}
	if _, ok := l.adapters[name]; ok {

		return errors.New("Adapter already exists")
	}

	tmpAdapters := make(map[string]adapter.Adapter, len(l.adapters)+1)

	for k, v := range l.adapters {
		tmpAdapters[k] = v
	}

	tmpAdapters[name] = adp
	l.adapters = tmpAdapters

	return nil
}

// RemoveAdapter remove a log adapter with give name key, return error in name not exists
func (l *gollog) RemoveAdapter(name string) error {
	if _, ok := l.adapters[name]; !ok {
		return errors.New("Adapter not exists")
	}

	tmpAdapters := make(map[string]adapter.Adapter, len(l.adapters)-1)

	for k, v := range l.adapters {
		if k != name {
			tmpAdapters[k] = v
		}
	}

	l.adapters = tmpAdapters

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
