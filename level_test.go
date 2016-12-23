package gol

import (
	"testing"
)

func TestLogLevel_String(t *testing.T) {
	tests := []struct {
		name  string
		level LogLevel
		want  string
	}{
		{
			"case1",
			DEBUG,
			"[DEBUG]",
		},
		{
			"case2",
			ALL,
			"[ALL]",
		},
		{
			"case3",
			INFO,
			"[INFO]",
		},
		{
			"case4",
			WARN,
			"[WARN]",
		},
		{
			"case5",
			ERROR,
			"[ERROR]",
		},
		{
			"case6",
			CRITICAL,
			"[CRITICAL]",
		},
		{
			"case7",
			LogLevel(10),
			"[UNKNOWN LOG LEVEL]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.level.String(); got != tt.want {
				t.Errorf("LogLevel.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLogLevel_ColorString(t *testing.T) {
	tests := []struct {
		name  string
		level LogLevel
		want  string
	}{
		{
			"case1",
			DEBUG,
			"\033[32m",
		},
		{
			"case2",
			ALL,
			"\033[0m",
		},
		{
			"case3",
			INFO,
			"\033[34m",
		},
		{
			"case4",
			WARN,
			"\033[33m",
		},
		{
			"case5",
			ERROR,
			"\033[31m",
		},
		{
			"case6",
			CRITICAL,
			"\033[35m",
		},
		{
			"case7",
			LogLevel(10),
			"\033[0m",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.level.ColorString(); got != tt.want {
				t.Errorf("LogLevel.ColorString() = %v, want %v", got, tt.want)
			}
		})
	}
}
