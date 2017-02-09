package fake

import (
	"time"

	"github.com/philchia/gol/adapter"
	"github.com/philchia/gol/level"
)

var _ adapter.Adapter = (*Writer)(nil)

// Writer is a fake adapter
// Use for test purpose
type Writer struct {
	withErr  error
	b        []byte
	logLevel level.LogLevel
}

// Write append bytes to b
func (w *Writer) Write(b []byte) (int, error) {
	if w.withErr != nil {
		return 0, w.withErr
	}
	time.Sleep(time.Nanosecond * 10)
	w.b = append(w.b, b...)
	return len(b), nil
}

// Close will do nothing
func (w *Writer) Close() error {
	if w.withErr != nil {
		return w.withErr
	}
	w.b = w.b[0:0]
	return nil
}

// NewAdapter create a fake adapter
func NewAdapter(l ...level.LogLevel) *Writer {
	w := new(Writer)
	if len(l) > 0 {
		w.logLevel = l[0]
	}
	return w
}

func (w *Writer) Level() level.LogLevel {
	return w.logLevel
}
