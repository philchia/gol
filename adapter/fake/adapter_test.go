package fake

import (
	"errors"
	"reflect"
	"testing"
)

func TestWriter_Write(t *testing.T) {
	type fields struct {
		withErr error
		b       []byte
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
				nil,
				[]byte{},
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
				errors.New("Hi"),
				[]byte{},
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
			w := &Writer{
				withErr: tt.fields.withErr,
				b:       tt.fields.b,
			}
			got, err := w.Write(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Writer.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Writer.Write() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriter_Close(t *testing.T) {
	type fields struct {
		withErr error
		b       []byte
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			"case1",
			fields{
				nil,
				nil,
			},
			false,
		},
		{
			"case2",
			fields{
				errors.New("hi"),
				nil,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Writer{
				withErr: tt.fields.withErr,
				b:       tt.fields.b,
			}
			if err := w.Close(); (err != nil) != tt.wantErr {
				t.Errorf("Writer.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewAdapter(t *testing.T) {
	tests := []struct {
		name string
		want *Writer
	}{
		{
			"case1",
			new(Writer),
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
