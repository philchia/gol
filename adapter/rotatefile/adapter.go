package rotatefile

import (
	"io"
	"os"
	"path/filepath"

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
	maxFileBackups     int
	maxByteSizePerFile ByteSize
	size               ByteSize
	fileName           string
	io.WriteCloser
}

// NewAdapter create a new rotate file adapter
func NewAdapter(name string, maxFileBackups int, maxBytesPerFile ByteSize) adapter.Adapter {
	path, err := filepath.Abs(name)
	if err != nil {
		return nil
	}
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil
	}
	info, err := os.Stat(path)
	if err != nil {
		return nil
	}

	adapter := &rotatefileAdapter{
		maxFileBackups:     maxFileBackups,
		maxByteSizePerFile: maxBytesPerFile,
		fileName:           name,
		WriteCloser:        file,
		size:               ByteSize(info.Size()),
	}
	return adapter
}

// Write implement Writer
func (r *rotatefileAdapter) Write(b []byte) (n int, err error) {
	n, err = r.WriteCloser.Write(b)
	r.size += ByteSize(n)
	return
}
