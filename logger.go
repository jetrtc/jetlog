package jetlog

import (
	"fmt"
	"log"
)

type LogLevel int

const (
	LogDebug = iota
	LogInfo
	LogWarning
	LogError
	LogFatal
)

type Logger interface {
	Fatalf(format string, v ...interface{})
	Debugf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warningf(format string, v ...interface{})
}

func NewDefaultLogger(logger *log.Logger) *DefaultLogger {
	return &DefaultLogger{logger: logger, Calldepth: 2}
}

type DefaultLogger struct {
	logger    *log.Logger
	Calldepth int
	Level     LogLevel
}

func (l *DefaultLogger) output(level LogLevel, format string, v ...interface{}) {
	if level < l.Level {
		return
	}
	label := ""
	switch level {
	case LogDebug:
		label = "D"
	case LogInfo:
		label = "I"
	case LogWarning:
		label = "W"
	case LogError:
		label = "E"
	case LogFatal:
		label = "F"
	}
	l.logger.Output(l.Calldepth, fmt.Sprintf("[%s] %s", label, fmt.Sprintf(format, v...)))
}

func (l *DefaultLogger) Fatalf(format string, v ...interface{}) {
	l.output(LogFatal, format, v...)
}

func (l *DefaultLogger) Debugf(format string, v ...interface{}) {
	l.output(LogDebug, format, v...)
}

func (l *DefaultLogger) Errorf(format string, v ...interface{}) {
	l.output(LogError, format, v...)
}

func (l *DefaultLogger) Infof(format string, v ...interface{}) {
	l.output(LogInfo, format, v...)
}

func (l *DefaultLogger) Warningf(format string, v ...interface{}) {
	l.output(LogWarning, format, v...)
}

type Loggable struct {
	Logger Logger
}

func (l Loggable) Fatalf(format string, v ...interface{}) {
	l.Logger.Fatalf(format, v...)
}

func (l Loggable) Debugf(format string, v ...interface{}) {
	l.Logger.Debugf(format, v...)
}

func (l Loggable) Errorf(format string, v ...interface{}) {
	l.Logger.Errorf(format, v...)
}

func (l Loggable) Infof(format string, v ...interface{}) {
	l.Logger.Infof(format, v...)
}

func (l Loggable) Warningf(format string, v ...interface{}) {
	l.Logger.Warningf(format, v...)
}
