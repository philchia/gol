package console

import (
	"os"

	"io"

	"github.com/philchia/gol/adapter"
	"github.com/philchia/gol/level"
)

var _ adapter.Adapter = (*consoleAdapter)(nil)

type consoleAdapter struct {
	io.WriteCloser
	logLevel level.LogLevel
}

// NewAdapter create a console adapter
func NewAdapter(l ...level.LogLevel) adapter.Adapter {
	c := &consoleAdapter{
		WriteCloser: os.Stderr,
	}
	if len(l) > 0 {
		c.logLevel = l[0]
	}
	return c
}

func (c *consoleAdapter) Level() level.LogLevel {
	return c.logLevel
}
