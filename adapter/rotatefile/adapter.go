package file

import (
	"io"
	"os"

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

type rotatefileAdapter struct {
	maxFileNum         int
	maxByteSizePerFile ByteSize
	fileName           string
	io.WriteCloser
}

func NewAdapter(name string, maxFileNum int, maxBytesPerFile ByteSize) adapter.Adapter {
	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil
	}

	adapter := &rotatefileAdapter{
		maxFileNum:         maxFileNum,
		maxByteSizePerFile: maxBytesPerFile,
		fileName:           name,
		WriteCloser:        file,
	}
	return adapter
}
