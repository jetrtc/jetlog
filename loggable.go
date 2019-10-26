package log

type Loggable struct {
	Logger Logger
}

func NewLoggable(logger Logger) *Loggable {
	return &Loggable{Logger: logger}
}

func (l *Loggable) Output(level LogLevel, format string, v ...interface{}) {
	l.Logger.Output(level, format, v...)
}

func (l *Loggable) Fatalf(format string, v ...interface{}) {
	l.Logger.Output(LogFatal, format, v...)
}

func (l *Loggable) Debugf(format string, v ...interface{}) {
	l.Logger.Output(LogDebug, format, v...)
}

func (l *Loggable) Errorf(format string, v ...interface{}) {
	l.Logger.Output(LogError, format, v...)
}

func (l *Loggable) Infof(format string, v ...interface{}) {
	l.Logger.Output(LogInfo, format, v...)
}

func (l *Loggable) Warningf(format string, v ...interface{}) {
	l.Logger.Output(LogWarning, format, v...)
}
