package gol

import (
	"github.com/philchia/gol/internal"
)

// LogLevel represent the level of the logger
type LogLevel int8

const (
	// ALL level log all log level
	ALL LogLevel = iota
	// DEBUG level log all log having level greater than or equal to DEBUG
	DEBUG
	// INFO level log all log having level greater than or equal to INFO
	INFO
	// WARN level log all log having level greater than or equal to WARN
	WARN
	// ERROR level log all log having level greater than or equal to ERROR
	ERROR
	// CRITICAL level log all log having level greater than or equal to CRITICAL
	CRITICAL
)

var levelMap = map[LogLevel][]byte{
	ALL:      []byte("\033[0m[ALL]"),
	DEBUG:    []byte("\033[32m[DEBUG]\033[0m"),
	INFO:     []byte("\033[34m[INFO]\033[0m"),
	WARN:     []byte("\033[33m[WARN]\033[0m"),
	ERROR:    []byte("\033[31m[ERROR]\033[0m"),
	CRITICAL: []byte("\033[35m[CRITICAL]\033[0m"),
}

// Bytes return the level string
func (level LogLevel) Bytes() []byte {
	if ret, ok := levelMap[level]; ok {
		return ret
	}
	return internal.Str2bytes("\033[0m[UNKNOWN LOG LEVEL]")
}
