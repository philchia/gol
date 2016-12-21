package gol

type LogLevel int8

const (
	ALL LogLevel = iota
	DEBUG
	INFO
	WARN
	ERROR
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
	ALL:      "ALL",
	DEBUG:    "DEBUG",
	INFO:     "INFO",
	WARN:     "WARN",
	ERROR:    "ERROR",
	CRITICAL: "CRITICAL",
}

func (level LogLevel) ColorString() string {
	if ret, ok := colorMap[level]; ok {
		return ret
	}
	return colorMap[ALL]
}

func (level LogLevel) String() string {
	if ret, ok := levelMap[level]; ok {
		return ret
	}
	return "UNKNOWN LOG LEVEL"
}
