package file

import (
	"io"
	"os"
	"reflect"
	"testing"

	"github.com/philchia/gol/adapter"
)

func TestNewFileAdapter(t *testing.T) {
	type args struct {
		pathToFile string
	}
	tests := []struct {
		name string
		args args
		want adapter.Adapter
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFileAdapter(tt.args.pathToFile); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFileAdapter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewConsoleAdapter(t *testing.T) {
	tests := []struct {
		name string
		want adapter.Adapter
	}{
		{
			"case1",
			&fileAdapter{
				writer: os.Stderr,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConsoleAdapter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConsoleAdapter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fileAdapter_Write(t *testing.T) {
	type fields struct {
		writer io.Writer
	}
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &fileAdapter{
				writer: tt.fields.writer,
			}
			if err := a.Write(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("fileAdapter.Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
