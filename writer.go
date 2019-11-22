package log

type Writer func(level Level, payload interface{})
