package log

import "context"

type Context interface {
	context.Context
	Sugar
}

func NewContext(ctx context.Context, sugar Sugar) Context {
	return &context_{Context: ctx, Sugar: sugar}
}

type context_ struct {
	context.Context
	Sugar
}
