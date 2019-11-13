package log

import "context"

type Context interface {
	context.Context
	Loggable
}

func NewContext(ctx context.Context, logger Logger) Context {
	return &context_{Context: ctx, Loggable: NewLoggable(logger)}
}

type context_ struct {
	context.Context
	Loggable
}
