package console

import (
	"os"

	"io"

	"github.com/philchia/gol/adapter"
)

var _ adapter.Adapter = (*consoleAdapter)(nil)

type consoleAdapter struct {
	io.WriteCloser
}

// NewAdapter create a console adapter
func NewAdapter() adapter.Adapter {
	return &consoleAdapter{
		os.Stderr,
	}
}
