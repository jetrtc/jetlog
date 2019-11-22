package log

type Processor func(v interface{}) interface{}

type Logger interface {
	Write(lv Level, payload interface{})
	Debug(payload interface{})
	Info(payload interface{})
	Warning(payload interface{})
	Error(payload interface{})
	Fatal(payload interface{})
}

type logger struct {
	processor Processor
	writer    Writer
}

func NewLogger(writer Writer) Logger {
	return &logger{writer: writer}
}

func (l *logger) Write(lv Level, payload interface{}) {
	l.writer(lv, payload)
}

func (l *logger) Debug(payload interface{}) {
	l.writer(Debug, payload)
}

func (l *logger) Info(payload interface{}) {
	l.writer(Info, payload)
}

func (l *logger) Warning(payload interface{}) {
	l.writer(Warning, payload)
}

func (l *logger) Error(payload interface{}) {
	l.writer(Error, payload)
}

func (l *logger) Fatal(payload interface{}) {
	l.writer(Fatal, payload)
}
