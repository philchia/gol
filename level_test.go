package gol

import "testing"

func TestLogLevel_String(t *testing.T) {
	tests := []struct {
		name  string
		level LogLevel
		want  string
	}{
		{
			"case1",
			DEBUG,
			"DEBUG",
		},
		{
			"case2",
			ALL,
			"ALL",
		},
		{
			"case3",
			INFO,
			"INFO",
		},
		{
			"case4",
			WARN,
			"WARN",
		},
		{
			"case5",
			ERROR,
			"ERROR",
		},
		{
			"case6",
			CRITICAL,
			"CRITICAL",
		},
		{
			"case7",
			LogLevel(10),
			"UNKNOWN LOG LEVEL",
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
