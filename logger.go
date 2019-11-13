package log

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
	Output(level LogLevel, format string, v ...interface{})
}

func NewDefaultLogger(logger *log.Logger) *DefaultLogger {
	l := &DefaultLogger{logger: logger, Calldepth: 2}
	l.logger = logger
	return l
}

type DefaultLogger struct {
	logger    *log.Logger
	Calldepth int
	Level     LogLevel
}

func (l *DefaultLogger) Output(level LogLevel, format string, v ...interface{}) {
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
