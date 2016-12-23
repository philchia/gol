package fake

import "time"

// Writer is a fake adapter
// Use for test purpose
type Writer struct {
	withErr error
	b       []byte
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
	return nil
}

// NewAdapter create a fake adapter
func NewAdapter() *Writer {
	return new(Writer)
}
