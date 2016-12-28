package gol

import (
	"bytes"
	"runtime"
	"time"
)

// SetOption set the shared logger's log options used to format log headerr
func (l *gollog) SetOption(option LogOption) {
	l.option = option
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

	writeDateTimeHeader(buf, l.option)
	writeFineLineHeader(buf, l.option, callDepth+1)
}

func writeDateTimeHeader(buf *bytes.Buffer, option LogOption) {
	if option&(Ldate|Ltime|Lmicroseconds) != 0 {
		var t = time.Now()
		if option&LUTC != 0 {
			t = t.UTC()
		}

		if option&Ldate != 0 {
			year, month, day := t.Date()
			itoa(buf, year, 4)
			buf.WriteByte('/')
			itoa(buf, int(month), 2)
			buf.WriteByte('/')
			itoa(buf, day, 2)
			buf.WriteByte(' ')
		}
		if option&(Ltime) != 0 {
			hour, min, sec := t.Clock()
			itoa(buf, hour, 2)
			buf.WriteByte(':')
			itoa(buf, min, 2)
			buf.WriteByte(':')
			itoa(buf, sec, 2)
		}
		if option&Lmicroseconds != 0 {
			buf.WriteByte('.')
			itoa(buf, t.Nanosecond()/1e3, 6)
		}
		buf.WriteByte(' ')
	}
}

func writeFineLineHeader(buf *bytes.Buffer, option LogOption, callDepth int) {
	if option&(Lshortfile|Llongfile) != 0 {
		file, line := getCaller(callDepth + 1)
		if option&Lshortfile != 0 {
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

func getCaller(depth int) (file string, line int) {

	var ok bool
	_, file, line, ok = runtime.Caller(depth)
	if !ok {
		file = "???"
		line = 0
	}
	return
}

func (l *gollog) generateLog(callDepth int, level LogLevel, msg string) *bytes.Buffer {
	buf := bufferPoolGet()
	l.generatePrefix(buf, callDepth)

	buf.Write(level.ColorString())
	buf.Write(level.String())
	buf.Write(ALL.ColorString())
	buf.WriteByte(' ')
	buf.WriteString(msg)
	buf.WriteByte('\n')
	return buf
}
