package gol

import (
	"github.com/philchia/gol/level"
)

// SetLevel set the shared logger's log level
func (l *gollog) SetLevel(level level.LogLevel) {
	l.level = level
}
