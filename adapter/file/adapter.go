package file

import (
	"os"

	"io"

	"path/filepath"

	"github.com/philchia/gol/adapter"
	"github.com/philchia/gol/level"
)

var _ adapter.Adapter = (*fileAdapter)(nil)

type fileAdapter struct {
	io.WriteCloser
	logLevel level.LogLevel
}

// NewAdapter create a file adapter with given file name, will automatically create a file if not exists
func NewAdapter(name string, l ...level.LogLevel) adapter.Adapter {
	path, err := filepath.Abs(name)
	if err != nil {
		return nil
	}
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil
	}

	adapter := &fileAdapter{
		WriteCloser: file,
	}
	if len(l) > 0 {
		adapter.logLevel = l[0]
	}
	return adapter
}

func (f *fileAdapter) Level() level.LogLevel {
	return f.logLevel
}
