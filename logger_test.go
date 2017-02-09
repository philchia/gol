package gol

import (
	"testing"

	"github.com/philchia/gol/level"
)

func TestNewLogger(t *testing.T) {
	type args struct {
		level level.LogLevel
	}
	tests := []struct {
		name string
		args args
		want Logger
	}{
		{
			"case1",
			args{
				level.DEBUG,
			},
			&gollog{
				level:  level.DEBUG,
				option: LstdFlags,
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
