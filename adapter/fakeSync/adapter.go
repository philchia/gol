package fakeSync

import "sync"

// ReadWriter is a fake adapter with sync read method
// Use for test purpose
type ReadWriter struct {
	withErr error
	wg      sync.WaitGroup
	b       []byte
}

// Write append bytes to b
func (w *ReadWriter) Write(b []byte) (int, error) {
	defer w.wg.Done()
	if w.withErr != nil {
		return 0, w.withErr
	}
	w.b = b
	return len(b), nil
}

// Close will do nothing
func (w *ReadWriter) Close() error {
	w.b = w.b[0:0]
	return nil
}

func (w *ReadWriter) Read() []byte {
	w.wg.Add(1)
	w.wg.Wait()
	return w.b
}

// NewAdapter create a fake adapter with sync read method
func NewAdapter() *ReadWriter {
	return new(ReadWriter)
}
