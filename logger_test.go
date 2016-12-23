package gol

import (
	"bytes"
	"testing"
)

func TestNewLogger(t *testing.T) {
	type args struct {
		level LogLevel
	}
	tests := []struct {
		name string
		args args
		want Logger
	}{
		{
			"case1",
			args{
				DEBUG,
			},
			&gollog{
				level:   DEBUG,
				option:  LstdFlags,
				logChan: make(chan *bytes.Buffer, 1024),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLogger(tt.args.level); got.(*gollog).level != tt.args.level || got.(*gollog).option != LstdFlags || got.(*gollog).logChan == nil {
				t.Errorf("NewLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}
