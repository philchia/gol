package console

import (
	"io"
	"reflect"
	"testing"

	"os"

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

func TestNewAdapter(t *testing.T) {
	tests := []struct {
		name string
		want adapter.Adapter
	}{
		{
			"case1",
			&consoleAdapter{
				writer: os.Stderr,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAdapter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAdapter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_consoleAdapter_Write(t *testing.T) {
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
		want    int
		wantErr bool
	}{
		{
			"case1",
			fields{
				writer: new(fakeWriter),
			},
			args{
				[]byte("Hello"),
			},
			5,
			false,
		},
		{
			"case2",
			fields{
				writer: &fakeWriter{
					withErr: errors.New("failed to write"),
				},
			},
			args{
				[]byte("Hello"),
			},
			0,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &consoleAdapter{
				writer: tt.fields.writer,
			}
			got, err := c.Write(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("consoleAdapter.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("consoleAdapter.Write() = %v, want %v", got, tt.want)
			}
		})
	}
}
