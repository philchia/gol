package gol

var logger Logger

func Debug(i ...interface{}) {
	logger.Debug(i...)
}

func Debugf(format string, i ...interface{}) {
	logger.Debugf(format, i...)
}

func Info(i ...interface{}) {
	logger.Info(i...)
}

func Infof(format string, i ...interface{}) {
	logger.Infof(format, i...)
}

func Warn(i ...interface{}) {
	logger.Warn(i...)
}

func Warnf(format string, i ...interface{}) {
	logger.Warnf(format, i...)
}

func Error(i ...interface{}) {
	logger.Error(i...)
}

func Errorf(format string, i ...interface{}) {
	logger.Errorf(format, i...)
}

func Critical(i ...interface{}) {
	logger.Critical(i...)
}

func Criticalf(format string, i ...interface{}) {
	logger.Criticalf(format, i...)
}
