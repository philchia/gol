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

var colorMap = map[LogLevel][]byte{
	ALL:      []byte("\033[0m"),
	DEBUG:    []byte("\033[32m"),
	INFO:     []byte("\033[34m"),
	WARN:     []byte("\033[33m"),
	ERROR:    []byte("\033[31m"),
	CRITICAL: []byte("\033[35m"),
}

var levelMap = map[LogLevel][]byte{
	ALL:      []byte("[ALL]"),
	DEBUG:    []byte("[DEBUG]"),
	INFO:     []byte("[INFO]"),
	WARN:     []byte("[WARN]"),
	ERROR:    []byte("[ERROR]"),
	CRITICAL: []byte("[CRITICAL]"),
}

// ColorString return the ascii color code represent the given log level
func (level LogLevel) ColorString() []byte {
	if ret, ok := colorMap[level]; ok {
		return ret
	}
	return colorMap[ALL]
}

// String return the level string
func (level LogLevel) String() []byte {
	if ret, ok := levelMap[level]; ok {
		return ret
	}
	return internal.Str2bytes("[UNKNOWN LOG LEVEL]")
}
