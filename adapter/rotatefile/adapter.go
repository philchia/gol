package rotatefile

import (
	"io"
	"log"
	"os"

	"errors"

	"strconv"

	"github.com/philchia/gol/adapter"
	"github.com/philchia/gol/internal"
)

// ByteSize represent file size in byte
type ByteSize int64

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
	adapter := &rotatefileAdapter{
		maxFileBackups:     maxFileBackups,
		maxByteSizePerFile: maxBytesPerFile,
		fileName:           name,
		size:               0,
	}
	if err := adapter.rotate(); err != nil {
		log.Println(err)
		return nil
	}
	return adapter
}

// Write implement Writer
func (r *rotatefileAdapter) Write(b []byte) (n int, err error) {
	writeLen := len(b)
	if ByteSize(writeLen) > r.maxByteSizePerFile {
		return 0, errors.New("write length exceeds maximum file size")
	}

	if r.size+ByteSize(writeLen) > r.maxByteSizePerFile {
		//rotate here
		if err := r.rotate(); err != nil {
			return 0, err
		}
	}

	n, err = r.WriteCloser.Write(b)
	r.size += ByteSize(n)
	return
}

func (r *rotatefileAdapter) rotate() error {
	// close file first
	if r.WriteCloser != nil {
		log.Println("not nil")
		if err := r.Close(); err != nil {
			log.Println("close err")
			return err
		}
	}
	// find next file
	for i := 1; i <= r.maxFileBackups; i++ {
		fileName := internal.JoinStrings(r.fileName, ".", strconv.FormatInt(int64(i), 10))
		info, err := os.Stat(fileName)
		if (err != nil && os.IsNotExist(err)) || (err == nil && ByteSize(info.Size()) < r.maxByteSizePerFile) {
			// open or create file here
			wc, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
			if err != nil {
				log.Println(err)
				return err
			}
			if info == nil {
				r.size = 0
			} else {
				r.size = ByteSize(info.Size())
			}
			r.WriteCloser = wc
			return nil
		}
	}

	// rotate file
	for i := 2; i <= r.maxFileBackups; i++ {
		fileName := internal.JoinStrings(r.fileName, ".", strconv.FormatInt(int64(i), 10))
		newFileName := internal.JoinStrings(r.fileName, ".", strconv.FormatInt(int64(i-1), 10))
		err := os.Rename(fileName, newFileName)
		if err != nil {
			return err
		}
	}

	// create a new file
	fileName := internal.JoinStrings(r.fileName, ".", strconv.FormatInt(int64(r.maxFileBackups), 10))
	wc, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	r.WriteCloser = wc
	r.size = 0
	return nil

}
