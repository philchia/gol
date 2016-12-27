package fakeSync

import (
	"errors"
	"reflect"
	"testing"
)

func TestReadWriter_Write(t *testing.T) {
	type fields struct {
		withErr error
		done    chan struct{}
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
				make(chan struct{}, 100),
				nil,
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
				make(chan struct{}, 100),
				nil,
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
			w := &ReadWriter{
				withErr: tt.fields.withErr,
				done:    tt.fields.done,
				b:       tt.fields.b,
			}
			got, err := w.Write(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadWriter.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadWriter.Write() = %v, want %v", got, tt.want)
			}
			if err == nil {
				read := w.Read()
				if !reflect.DeepEqual(read, tt.args.b) {
					t.Errorf("ReadWriter.Read() = %v, want %v", read, tt.args.b)
				}
			}
		})
	}
}

func TestReadWriter_Close(t *testing.T) {
	type fields struct {
		withErr error
		done    chan struct{}
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
				errors.New("Hi"),
				make(chan struct{}, 100),
				nil,
			},
			true,
		},
		{
			"case2",
			fields{
				nil,
				make(chan struct{}, 100),
				nil,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &ReadWriter{
				withErr: tt.fields.withErr,
				done:    tt.fields.done,
				b:       tt.fields.b,
			}
			if err := w.Close(); (err != nil) != tt.wantErr {
				t.Errorf("ReadWriter.Close() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewAdapter(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"case1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAdapter(); got == nil {
				t.Errorf("NewAdapter() = nil, want %v", got)
			}
		})
	}
}
