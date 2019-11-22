package log

import (
	"fmt"
	"io"
	"log"
)

const (
	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	Ltime                         // the time in the local time zone: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
)

func GoLogger(level Level, out io.Writer, prefix string, flag int) Writer {
	calldepth := 2
	logger := log.New(out, prefix, flag)
	return func(lv Level, v interface{}) {
		if lv < level {
			return
		}
		var msg string
		switch v := v.(type) {
		case string:
			msg = fmt.Sprintf("[%s] %s", lv.Tag(), v)
		default:
			msg = fmt.Sprintf("[%s] %v", lv.Tag(), v)
		}
		logger.Output(calldepth, msg)
	}
}
