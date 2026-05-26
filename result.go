package pond

import (
	"context"
)

// ResultPool is a pool that can be used to submit tasks that return a result.
type ResultPool[R any] interface {
	BasePool

	// Submits a task to the pool and returns a future that can be used to wait for the task to complete and get the result.
	// The pool will not accept new tasks after it has been stopped.
	// If the pool has been stopped, this method will return ErrPoolStopped.
	Submit(task func() R) ResultTask[R]

	// Submits a task to the pool and returns a future that can be used to wait for the task to complete and get the result.
	// The task function must return a result and an error.
	// The pool will not accept new tasks after it has been stopped.
	// If the pool has been stopped, this method will return ErrPoolStopped.
	SubmitErr(task func() (R, error)) ResultTask[R]

	// Attempts to submit a task to the pool and returns a future that can be used to wait for the task to complete
	// and a boolean indicating whether the task was submitted successfully.
	// The pool will not accept new tasks after it has been stopped.
	// If the pool has been stopped, this method will return ErrPoolStopped.
	TrySubmit(task func() R) (ResultTask[R], bool)

	// Attempts to submit a task to the pool and returns a future that can be used to wait for the task to complete
	// and a boolean indicating whether the task was submitted successfully.
	// The task function must return a result and an error.
	// The pool will not accept new tasks after it has been stopped.
	// If the pool has been stopped, this method will return ErrPoolStopped.
	TrySubmitErr(task func() (R, error)) (ResultTask[R], bool)

	// Creates a new subpool with the specified maximum concurrency and options.
	NewSubpool(maxConcurrency int, options ...Option) ResultPool[R]

	// Creates a new task group.
	NewGroup() ResultTaskGroup[R]

	// Creates a new task group with the specified context.
	NewGroupContext(ctx context.Context) ResultTaskGroup[R]
}

type resultPool[R any] struct {
	*pool
}

func (p *resultPool[R]) NewGroup() ResultTaskGroup[R] { _ = "STUB: not implemented"; return nil }

func (p *resultPool[R]) NewGroupContext(ctx context.Context) ResultTaskGroup[R] {
	_ = "STUB: not implemented"
	return nil
}

func (p *resultPool[R]) Submit(task func() R) ResultTask[R] { _ = "STUB: not implemented"; return nil }

func (p *resultPool[R]) SubmitErr(task func() (R, error)) ResultTask[R] {
	_ = "STUB: not implemented"
	return nil
}

func (p *resultPool[R]) TrySubmit(task func() R) (ResultTask[R], bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (p *resultPool[R]) TrySubmitErr(task func() (R, error)) (ResultTask[R], bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (p *resultPool[R]) submit(task any, nonBlocking bool) (ResultTask[R], bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (p *resultPool[R]) NewSubpool(maxConcurrency int, options ...Option) ResultPool[R] {
	_ = "STUB: not implemented"
	return nil
}

func newResultPool[R any](maxConcurrency int, parent *pool, options ...Option) *resultPool[R] {
	_ = "STUB: not implemented"
	return nil
}

// NewResultPool creates a new result pool with the given maximum concurrency and options.
// Result pools are generic pools that can be used to submit tasks that return a result.
// The new maximum concurrency must be greater than or equal to 0 (0 means no limit).
func NewResultPool[R any](maxConcurrency int, options ...Option) ResultPool[R] {
	_ = "STUB: not implemented"
	return nil
}
