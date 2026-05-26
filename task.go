package pond

import (
	"context"
	"errors"
)

var ErrPanic = errors.New("task panicked")

var ErrContextCanceled = errors.New("context canceled")

type wrappedTask[R any, C func(error) | func(R, error)] struct {
	task          any
	callback      C
	ctx           context.Context
	panicRecovery bool
}

func (t wrappedTask[R, C]) Run() error { _ = "STUB: not implemented"; return nil }

func wrapTask[R any, C func(error) | func(R, error)](task any, callback C, ctx context.Context, panicRecovery bool) func() error {
	_ = "STUB: not implemented"
	return nil
}

func invokeTask[R any](task any, panicRecovery bool) (output R, err error) {
	_ = "STUB: not implemented"
	return *new(R), nil
}
