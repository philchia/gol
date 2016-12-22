package console

import (
	"os"

	"io"

	"github.com/philchia/gol/adapter"
)

type consoleAdapter struct {
	writer io.Writer
}

func NewAdapter() adapter.Adapter {
	return &consoleAdapter{
		writer: os.Stderr,
	}
}

func (c *consoleAdapter) Write(b []byte) (int, error) {
	return c.writer.Write(b)
}
