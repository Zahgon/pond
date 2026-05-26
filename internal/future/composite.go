package future

import (
	"context"
	"sync"
)

type CompositeFutureResolver[V any] func(index int, value V, err error)

type compositeResolution[V any] struct {
	index int
	value V
}

type compositeErrorResolution struct {
	index int
	err   error
}

func (e *compositeErrorResolution) Error() string { _ = "STUB: not implemented"; return "" }

type waitListener struct {
	count int
	ch    chan struct{}
}

type CompositeFuture[V any] struct {
	ctx         context.Context
	cancel      context.CancelCauseFunc
	resolutions []compositeResolution[V]
	mutex       sync.Mutex
	listeners   []waitListener
}

func NewCompositeFuture[V any](ctx context.Context) (*CompositeFuture[V], CompositeFutureResolver[V]) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *CompositeFuture[V]) Done(count int) <-chan struct{} { _ = "STUB: not implemented"; return nil }

// Return immediately if the context is already canceled or the count is already reached

// Register a listener

func (f *CompositeFuture[V]) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (f *CompositeFuture[V]) Cancel(cause error) { _ = "STUB: not implemented"; return }

// Cancel the context

// Notify listeners

func (f *CompositeFuture[V]) Wait(count int) ([]V, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Register a listener

// Wait for the listener to be notified or the context to be canceled

func (f *CompositeFuture[V]) resolve(index int, value V, err error) {
	_ = "STUB: not implemented"
	return
}

// Cancel the context if an error occurred

// Save the resolution

// Notify listeners

func (f *CompositeFuture[V]) getResult(count int) (values []V, err error) {
	_ = "STUB: not implemented"
	return nil, nil

	// If we have enough results, return them
}

// Get sorted resolution values

// Unwrap the error resolution

// If the context is canceled and we have collected enough results, return nil error
// because we assume that context cancellation happened after the last resolution.

func (f *CompositeFuture[V]) notifyListeners() { _ = "STUB: not implemented"; return }

// Notify listeners
