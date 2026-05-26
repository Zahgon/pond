package future

import (
	"context"
)

type FutureResolver func(err error)

// A Future represents a value that will be available in the Future.
// It is always associated with a context that can be used to wait for the value to be available.
// When the parent context is canceled, the Future will be canceled as well.
type Future struct {
	ctx context.Context
}

func (f *Future) Done() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (f *Future) Err() error { _ = "STUB: not implemented"; return nil }

// Wait waits for the future to complete and returns any error that occurred.
func (f *Future) Wait() error { _ = "STUB: not implemented"; return nil }

func NewFuture(ctx context.Context) (*Future, FutureResolver) {
	_ = "STUB: not implemented"
	return nil, *new(FutureResolver)
}

type futureResolution struct {
	err error
}

func (v *futureResolution) Error() string { _ = "STUB: not implemented"; return "" }
