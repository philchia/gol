package gol

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

var colorMap = map[LogLevel]string{
	ALL:      "\033[0m",
	DEBUG:    "\033[32m",
	INFO:     "\033[34m",
	WARN:     "\033[33m",
	ERROR:    "\033[31m",
	CRITICAL: "\033[35m",
}

var levelMap = map[LogLevel]string{
	ALL:      "[ALL]",
	DEBUG:    "[DEBUG]",
	INFO:     "[INFO]",
	WARN:     "[WARN]",
	ERROR:    "[ERROR]",
	CRITICAL: "[CRITICAL]",
}

// ColorString return the ascii color code represent the given log level
func (level LogLevel) ColorString() string {
	if ret, ok := colorMap[level]; ok {
		return ret
	}
	return colorMap[ALL]
}

// String return the level string
func (level LogLevel) String() string {
	if ret, ok := levelMap[level]; ok {
		return ret
	}
	return "[UNKNOWN LOG LEVEL]"
}
