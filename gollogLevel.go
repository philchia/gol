package gol

// SetLevel set the shared logger's log level
func (l *gollog) SetLevel(level LogLevel) {
	l.level = level
}
