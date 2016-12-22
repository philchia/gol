package file

import (
	"io"
	"reflect"
	"testing"

	"errors"

	"github.com/philchia/gol/adapter"
)

type fakeWriter struct {
	withErr error
}

func (w *fakeWriter) Write(b []byte) (int, error) {
	if w.withErr != nil {
		return 0, w.withErr
	}
	return len(b), nil
}

func (w *fakeWriter) Close() error {
	return nil
}

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
			if got := NewAdapter(tt.args.pathToFile); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFileAdapter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fileAdapter_Write(t *testing.T) {
	type fields struct {
		writer io.WriteCloser
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
		{
			"case1",
			fields{
				new(fakeWriter),
			},
			args{
				[]byte("hi"),
			},
			false,
		},
		{
			"case2",
			fields{
				&fakeWriter{
					withErr: errors.New("test"),
				},
			},
			args{
				[]byte("hi"),
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &fileAdapter{
				tt.fields.writer,
			}
			if _, err := a.Write(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("fileAdapter.Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
