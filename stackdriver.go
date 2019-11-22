package log

import (
	"cloud.google.com/go/logging"
)

func StackDriver(logger *logging.Logger) Writer {
	return func(lv Level, v interface{}) {
		var severity logging.Severity
		switch lv {
		case Debug:
			severity = logging.Debug
		case Info:
			severity = logging.Info
		case Warning:
			severity = logging.Warning
		case Error:
			severity = logging.Error
		case Fatal:
			severity = logging.Critical
		}
		logger.Log(logging.Entry{Severity: severity, Payload: v})
	}
}
