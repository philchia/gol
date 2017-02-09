package fakeSync

import (
	"github.com/philchia/gol/adapter"
	"github.com/philchia/gol/level"
)

var _ adapter.Adapter = (*ReadWriter)(nil)

// ReadWriter is a fake adapter with sync read method
// Use for test purpose
type ReadWriter struct {
	withErr  error
	done     chan struct{}
	b        []byte
	logLevel level.LogLevel
}

// Write append bytes to b
func (w *ReadWriter) Write(b []byte) (int, error) {
	defer func() { w.done <- struct{}{} }()
	if w.withErr != nil {
		return 0, w.withErr
	}
	w.b = b
	return len(b), nil
}

// Close will do nothing
func (w *ReadWriter) Close() error {
	if w.withErr != nil {
		return w.withErr
	}
	w.b = w.b[0:0]
	return nil
}

func (w *ReadWriter) Read() []byte {
	<-w.done
	return w.b
}

// NewAdapter create a fake adapter with sync read method
func NewAdapter(l ...level.LogLevel) *ReadWriter {
	w := &ReadWriter{
		done: make(chan struct{}, 100),
	}
	if len(l) > 0 {
		w.logLevel = l[0]
	}
	return w
}

func (w *ReadWriter) Level() level.LogLevel {
	return w.logLevel
}
