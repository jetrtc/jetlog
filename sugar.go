package log

import "fmt"

type Sugar interface {
	Logger
	Writef(lv Level, format string, v ...interface{})
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warningf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Fatalf(format string, v ...interface{})
}

type sugar struct {
	Logger
}

func NewSugar(logger Logger) Sugar {
	return &sugar{Logger: logger}
}

func (l *sugar) Writef(lv Level, format string, v ...interface{}) {
	l.Write(lv, fmt.Sprintf(format, v...))
}

func (l *sugar) Debugf(format string, v ...interface{}) {
	l.Debug(fmt.Sprintf(format, v...))
}

func (l *sugar) Infof(format string, v ...interface{}) {
	l.Info(fmt.Sprintf(format, v...))
}

func (l *sugar) Warningf(format string, v ...interface{}) {
	l.Warning(fmt.Sprintf(format, v...))
}

func (l *sugar) Errorf(format string, v ...interface{}) {
	l.Error(fmt.Sprintf(format, v...))
}

func (l *sugar) Fatalf(format string, v ...interface{}) {
	l.Fatal(fmt.Sprintf(format, v...))
}
