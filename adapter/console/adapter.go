package console

import (
	"os"

	"io"

	"github.com/philchia/gol/adapter"
)

type consoleAdapter struct {
	io.WriteCloser
}

func NewAdapter() adapter.Adapter {
	return &consoleAdapter{
		os.Stderr,
	}
}
