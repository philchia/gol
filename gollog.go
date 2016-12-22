package gol

import (
	"fmt"
	"os"
	"runtime"

	"bytes"

	"time"

	"sync/atomic"

	"github.com/philchia/gol/adapter"
	"github.com/philchia/gol/internal"
)

type gollog struct {
	level       LogLevel
	option      LogOption
	adapters    []adapter.Adapter
	logChan     chan string
	doneChan    chan struct{}
	signalChan  chan os.Signal
	exitingFlag uint64
}

func (l *gollog) exiting() bool {
	return atomic.LoadUint64(&l.exitingFlag) == 1
}

func (l *gollog) setExiting(flag bool) {
	if flag {
		atomic.StoreUint64(&l.exitingFlag, 1)
	} else {
		atomic.StoreUint64(&l.exitingFlag, 0)
	}
}

func (l *gollog) msgPump() {

	for msg := range l.logChan {
		for _, adap := range l.adapters {
			adap.Write(internal.Str2bytes(msg))
		}
	}

	close(l.doneChan)
}

func (l *gollog) put(msg string) {
	if l.exiting() {
		return
	}
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

func (l *gollog) generatePrefix(callDepth int) string {
	var buf bytes.Buffer

	var t = time.Now()
	if l.option&LUTC != 0 {
		t = t.UTC()
	}
	if l.option&(Ldate|Ltime|Lmicroseconds) != 0 {
		if l.option&Ldate != 0 {
			year, month, day := t.Date()
			itoa(&buf, year, 4)
			buf.WriteByte('/')
			itoa(&buf, int(month), 2)
			buf.WriteByte('/')
			itoa(&buf, day, 2)
			buf.WriteByte(' ')
		}
		if l.option&(Ltime|Lmicroseconds) != 0 {
			hour, min, sec := t.Clock()
			itoa(&buf, hour, 2)
			buf.WriteByte(':')
			itoa(&buf, min, 2)
			buf.WriteByte(':')
			itoa(&buf, sec, 2)
			if l.option&Lmicroseconds != 0 {
				buf.WriteByte('.')
				itoa(&buf, t.Nanosecond()/1e3, 6)
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
		itoa(&buf, line, -1)
		buf.WriteString(": ")
	}

	return buf.String()
}

func (l *gollog) generateLog(callDepth int, level LogLevel, msg string) string {
	prefix := l.generatePrefix(callDepth)
	return internal.JoinStrings(prefix, level.ColorString(), "[", level.String(), "] ", ALL.ColorString(), msg, "\n")
}

func (l *gollog) Debug(i ...interface{}) {
	if l.level > DEBUG {
		return
	}
	msg := l.generateLog(2, DEBUG, fmt.Sprint(i...))
	l.put(msg)
}

func (l *gollog) Debugf(format string, i ...interface{}) {
	if l.level > DEBUG {
		return
	}
	msg := l.generateLog(2, DEBUG, fmt.Sprintf(format, i...))
	l.put(msg)
}

func (l *gollog) Info(i ...interface{}) {
	if l.level > INFO {
		return
	}
	msg := l.generateLog(2, INFO, fmt.Sprint(i...))
	l.put(msg)
}

func (l *gollog) Infof(format string, i ...interface{}) {
	if l.level > INFO {
		return
	}
	msg := l.generateLog(2, INFO, fmt.Sprintf(format, i...))
	l.put(msg)
}

func (l *gollog) Warn(i ...interface{}) {
	if l.level > WARN {
		return
	}
	msg := l.generateLog(2, WARN, fmt.Sprint(i...))
	l.put(msg)
}

func (l *gollog) Warnf(format string, i ...interface{}) {
	if l.level > WARN {
		return
	}
	msg := l.generateLog(2, WARN, fmt.Sprintf(format, i...))
	l.put(msg)
}

func (l *gollog) Error(i ...interface{}) {
	if l.level > ERROR {
		return
	}
	msg := l.generateLog(2, ERROR, fmt.Sprint(i...))
	l.put(msg)
}

func (l *gollog) Errorf(format string, i ...interface{}) {
	if l.level > ERROR {
		return
	}
	msg := l.generateLog(2, ERROR, fmt.Sprintf(format, i...))
	l.put(msg)
}

func (l *gollog) Critical(i ...interface{}) {
	if l.level > CRITICAL {
		return
	}
	msg := l.generateLog(2, CRITICAL, fmt.Sprint(i...))
	l.put(msg)
}

func (l *gollog) Criticalf(format string, i ...interface{}) {
	if l.level > CRITICAL {
		return
	}
	msg := l.generateLog(2, CRITICAL, fmt.Sprintf(format, i...))
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

func (l *gollog) flush() {
	l.setExiting(true)
	close(l.logChan)
	<-l.doneChan
}
