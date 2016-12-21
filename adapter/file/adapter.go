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

func NewAdapter(pathToFile string) adapter.Adapter {
	file, err := os.Open(pathToFile)
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
