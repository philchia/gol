package rotatefile

import (
	"io"
	"os"

	"github.com/philchia/gol/adapter"
)

// ByteSize represent file size in byte
type ByteSize float64

const (
	_ ByteSize = 1 << (iota * 10)
	// KB = 1 kb bytes
	KB
	// MB = 1 mb bytes
	MB
	// GB = 1 gb bytes
	GB
	// TB = 1 tb bytes
	TB
	// PB = 1 pb bytes
	PB
	// EB = 1 eb bytes
	EB
	// ZB = 1 zb bytes
	ZB
	// YB = 1 yb bytes
	YB
)

var _ adapter.Adapter = (*rotatefileAdapter)(nil)

type rotatefileAdapter struct {
	maxFileNum         int
	maxByteSizePerFile ByteSize
	fileName           string
	io.WriteCloser
}

// NewAdapter create a new rotate file adapter
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
