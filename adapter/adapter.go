package adapter

import (
	"io"
)

// Adapter write log to underly writer
type Adapter interface {
	io.Writer
}
