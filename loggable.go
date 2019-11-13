package log

type Loggable interface {
	Logger
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warningf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Fatalf(format string, v ...interface{})
}

func NewLoggable(logger Logger) Loggable {
	return &loggable{logger: logger}
}

type loggable struct {
	logger Logger
}

func (c *loggable) Output(level LogLevel, format string, v ...interface{}) {
	c.logger.Output(level, format, v...)
}

func (c *loggable) Debugf(format string, v ...interface{}) {
	c.logger.Output(LogDebug, format, v...)
}

func (c *loggable) Infof(format string, v ...interface{}) {
	c.logger.Output(LogInfo, format, v...)
}

func (c *loggable) Warningf(format string, v ...interface{}) {
	c.logger.Output(LogWarning, format, v...)
}

func (c *loggable) Errorf(format string, v ...interface{}) {
	c.logger.Output(LogError, format, v...)
}

func (c *loggable) Fatalf(format string, v ...interface{}) {
	c.logger.Output(LogFatal, format, v...)
}
