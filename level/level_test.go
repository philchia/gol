package level

import "testing"

func TestLogLevel_Bytes(t *testing.T) {
	tests := []struct {
		name  string
		level LogLevel
		want  string
	}{
		{
			"case1",
			DEBUG,
			"\033[32m[DEBUG]\033[0m",
		},
		{
			"case2",
			ALL,
			"\033[0m[ALL]",
		},
		{
			"case3",
			INFO,
			"\033[34m[INFO]\033[0m",
		},
		{
			"case4",
			WARN,
			"\033[33m[WARN]\033[0m",
		},
		{
			"case5",
			ERROR,
			"\033[31m[ERROR]\033[0m",
		},
		{
			"case6",
			CRITICAL,
			"\033[35m[CRITICAL]\033[0m",
		},
		{
			"case7",
			LogLevel(10),
			"\033[0m[UNKNOWN LOG LEVEL]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.level.Bytes(); string(got) != tt.want {
				t.Errorf("LogLevel.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
