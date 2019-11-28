package log

import (
	"fmt"
)

var (
	red      = color("\033[1;31m%s\033[0m")
	yellow   = color("\033[1;33m%s\033[0m")
	magenta  = color("\033[1;35m%s\033[0m")
	white    = color("\033[1;37m%s\033[0m")
	darkGray = color("\033[1;90m%s\033[0m")
)

func color(colorString string) func(msg string, args ...interface{}) string {
	sprint := func(msg string, args ...interface{}) string {
		return fmt.Sprintf(colorString, fmt.Sprintf(msg, args...))
	}
	return sprint
}

func Color(writer Writer) Writer {
	return func(lv Level, v interface{}) {
		switch v := v.(type) {
		case string:
			switch lv {
			case Debug:
				writer(lv, darkGray(v))
			case Info:
				writer(lv, white(v))
			case Warning:
				writer(lv, yellow(v))
			case Error:
				writer(lv, red(v))
			case Fatal:
				writer(lv, magenta(v))
			default:
				writer(lv, v)
			}
		default:
			writer(lv, v)
		}
	}
}
