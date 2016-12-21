package file

import (
	"io"
	"os"

	"github.com/philchia/gol/adapter"
)

type fileAdapter struct {
	writer io.Writer
}

func NewFileAdapter(pathToFile string) adapter.Adapter {
	file, err := os.Open(pathToFile)
	if err != nil {
		return nil
	}

	adapter := &fileAdapter{
		writer: file,
	}
	return adapter
}

func NewConsoleAdapter() adapter.Adapter {
	adapter := &fileAdapter{
		writer: os.Stderr,
	}
	return adapter
}

func (a *fileAdapter) Write(b []byte) error {
	_, err := a.writer.Write(b)
	return err
}
