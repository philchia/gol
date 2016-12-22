package file

import (
	"os"

	"io"

	"github.com/philchia/gol/adapter"
)

type ByteSize float64

const (
	_ ByteSize = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

type fileAdapter struct {
	rotateFile  bool
	rotateCount int
	rotateSize  ByteSize
	fileName    string
	file        io.Writer
}

func NewAdapter(name string) adapter.Adapter {
	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil
	}

	adapter := &fileAdapter{
		file: file,
	}
	return adapter
}

func (a *fileAdapter) Write(b []byte) (int, error) {
	return a.file.Write(b)
}
