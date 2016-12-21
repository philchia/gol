package console

import (
	"os"

	"github.com/philchia/gol/adapter"
)

type consoleAdapter struct {
}

func NewAdapter() adapter.Adapter {
	return &consoleAdapter{}
}

func (c *consoleAdapter) Write(b []byte) (int, error) {
	return os.Stderr.Write(b)
}
