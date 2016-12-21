package file

import (
	"os"

	"github.com/philchia/gol/adapter"
)

type fileAdapter struct {
	file *os.File
}

func NewFileAdapter(pathToFile string) adapter.Adapter {
	file, err := os.Open(pathToFile)
	if err != nil {
		return nil
	}

	adapter := &fileAdapter{
		file: file,
	}
	return adapter
}

func NewConsoleAdapter() adapter.Adapter {
	adapter := &fileAdapter{
		file: os.Stderr,
	}
	return adapter
}

func (a *fileAdapter) Write(b []byte) error {
	_, err := a.file.Write(b)
	return err
}
