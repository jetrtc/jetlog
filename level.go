package log

type Level int

const (
	Debug = iota
	Info
	Warning
	Error
	Fatal
)

func (lv Level) Tag() string {
	switch lv {
	case Debug:
		return "D"
	case Info:
		return "I"
	case Warning:
		return "W"
	case Error:
		return "E"
	case Fatal:
		return "F"
	}
	return "?"
}
