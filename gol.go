package gol

import "github.com/philchia/gol/adapter"
import "github.com/philchia/gol/level"

var logger Logger

func init() {
	logger = NewLogger(level.ALL)
}

// Debug will prinnt log as DEBUG level
func Debug(i ...interface{}) {
	logger.Debug(i...)
}

// Debugf will prinnt log as DEBUG level
func Debugf(format string, i ...interface{}) {
	logger.Debugf(format, i...)
}

// Info will prinnt log as INFO level
func Info(i ...interface{}) {
	logger.Info(i...)
}

// Infof will prinnt log as INFO level
func Infof(format string, i ...interface{}) {
	logger.Infof(format, i...)
}

// Warn will prinnt log as WARN level
func Warn(i ...interface{}) {
	logger.Warn(i...)
}

// Warnf will prinnt log as WARN level
func Warnf(format string, i ...interface{}) {
	logger.Warnf(format, i...)
}

// Error will prinnt log as ERROR level
func Error(i ...interface{}) {
	logger.Error(i...)
}

// Errorf will prinnt log as ERROR level
func Errorf(format string, i ...interface{}) {
	logger.Errorf(format, i...)
}

// Critical will prinnt log as CRITICAL level
func Critical(i ...interface{}) {
	logger.Critical(i...)
}

// Criticalf will prinnt log as CRITICAL level
func Criticalf(format string, i ...interface{}) {
	logger.Criticalf(format, i...)
}

// SetLevel set the shared logger's log level
func SetLevel(level level.LogLevel) {
	logger.SetLevel(level)
}

// SetOption set the shared logger's log options used to format log headerr
func SetOption(option LogOption) {
	logger.SetOption(option)
}

// AddLogAdapter add a log adapter which implement the adapter.Adapter interface with give name key, return error if name already exists
func AddLogAdapter(name string, adapter adapter.Adapter) error {
	return logger.AddLogAdapter(name, adapter)
}

// RemoveAdapter remove a log adapter with give name key, return error in name not exists
func RemoveAdapter(name string) error {
	return logger.RemoveAdapter(name)
}

// Flush flush all buffered log and call Close() on all adapters
func Flush() {
	logger.Flush()
}
