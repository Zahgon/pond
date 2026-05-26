package future

import (
	"context"
)

type ValueFutureResolver[V any] func(value V, err error)

// A Future represents a value that will be available in the Future.
// It is always associated with a context that can be used to wait for the value to be available.
// When the parent context is canceled, the Future will be canceled as well.
type ValueFuture[V any] struct {
	ctx context.Context
}

func (f *ValueFuture[V]) Done() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (f *ValueFuture[V]) Result() (V, error) { _ = "STUB: not implemented"; return *new(V), nil }

// Get waits for the future to complete and returns the output and any error that occurred.
func (f *ValueFuture[V]) Wait() (V, error) { _ = "STUB: not implemented"; return *new(V), nil }

func NewValueFuture[V any](ctx context.Context) (*ValueFuture[V], ValueFutureResolver[V]) {
	_ = "STUB: not implemented"
	return nil, nil
}

type valueFutureResolution[V any] struct {
	value V
	err   error
}

func (v *valueFutureResolution[V]) Error() string { _ = "STUB: not implemented"; return "" }
