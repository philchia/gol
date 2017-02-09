package adapter

import (
	"io"

	"github.com/philchia/gol/level"
)

// Adapter write log to underly writer
type Adapter interface {
	io.WriteCloser
	Level() level.LogLevel
}
