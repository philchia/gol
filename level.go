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

var levelMap = map[LogLevel]string{
	ALL:      "ALL",
	DEBUG:    "DEBUG",
	INFO:     "INFO",
	WARN:     "WARN",
	ERROR:    "ERROR",
	CRITICAL: "CRITICAL",
}

func (level LogLevel) String() string {
	if ret, ok := levelMap[level]; ok {
		return ret
	}
	return "UNKNOWN LOG LEVEL"
}
