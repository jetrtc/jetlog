package log

import (
	"context"
	"time"
)

type Context struct {
	Loggable
	Context context.Context
}

func NewContext(logger Logger, ctx context.Context) *Context {
	return &Context{Loggable: Loggable{Logger: logger}, Context: ctx}
}

func (ctx *Context) Deadline() (deadline time.Time, ok bool) {
	return ctx.Context.Deadline()
}

func (ctx *Context) Done() <-chan struct{} {
	return ctx.Context.Done()
}

func (ctx *Context) Err() error {
	return ctx.Context.Err()
}

func (ctx *Context) Value(key interface{}) interface{} {
	return ctx.Context.Value(key)
}
