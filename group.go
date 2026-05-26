package pond

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"

	"github.com/alitto/pond/v2/internal/future"
)

var ErrGroupStopped = errors.New("task group stopped")

// TaskGroup represents a group of tasks that can be executed concurrently.
// The group can be waited on to block until all tasks have completed.
// If any of the tasks return an error, the group will return the first error encountered.
type TaskGroup interface {

	// Submits a task to the group.
	Submit(tasks ...func()) TaskGroup

	// Submits a task to the group that can return an error.
	SubmitErr(tasks ...func() error) TaskGroup

	// Waits for all tasks in the group to complete.
	// If any of the tasks return an error, the group will return the first error encountered.
	// If the context is cancelled, the group will return the context error.
	// If the group is stopped, the group will return ErrGroupStopped.
	// If a task is running when the context is cancelled or the group is stopped, the task will be allowed to complete before returning.
	Wait() error

	// Returns a channel that is closed when all tasks in the group have completed, a task returns an error, or the group is stopped.
	Done() <-chan struct{}

	// Stops the group and cancels all remaining tasks. Running tasks are not interrupted.
	Stop()

	// Returns the context associated with this group.
	// This context will be cancelled when either the parent context is cancelled
	// or any task in the group returns an error, whichever comes first.
	Context() context.Context
}

// ResultTaskGroup represents a group of tasks that can be executed concurrently.
// As opposed to TaskGroup, the tasks in a ResultTaskGroup yield a result.
// The group can be waited on to block until all tasks have completed.
// If any of the tasks return an error, the group will return the first error encountered.
type ResultTaskGroup[O any] interface {

	// Submits a task to the group.
	Submit(tasks ...func() O) ResultTaskGroup[O]

	// Submits a task to the group that can return an error.
	SubmitErr(tasks ...func() (O, error)) ResultTaskGroup[O]

	// Waits for all tasks in the group to complete and returns the results of each task in the order they were submitted.
	// If any of the tasks return an error, the group will return the first error encountered.
	// If the context is cancelled, the group will return the context error.
	// If the group is stopped, the group will return ErrGroupStopped.
	// If a task is running when the context is cancelled or the group is stopped, the task will be allowed to complete before returning.
	Wait() ([]O, error)

	// Returns a channel that is closed when all tasks in the group have completed, a task returns an error, or the group is stopped.
	Done() <-chan struct{}

	// Stops the group and cancels all remaining tasks. Running tasks are not interrupted.
	Stop()
}

type result[O any] struct {
	Output O
	Err    error
}

type abstractTaskGroup[T func() | func() O, E func() error | func() (O, error), O any] struct {
	pool           *pool
	nextIndex      atomic.Int64
	taskWaitGroup  sync.WaitGroup
	future         *future.CompositeFuture[*result[O]]
	futureResolver future.CompositeFutureResolver[*result[O]]
}

func (g *abstractTaskGroup[T, E, O]) Done() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (g *abstractTaskGroup[T, E, O]) Stop() { _ = "STUB: not implemented"; return }

func (g *abstractTaskGroup[T, E, O]) Context() context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func (g *abstractTaskGroup[T, E, O]) Submit(tasks ...T) *abstractTaskGroup[T, E, O] {
	_ = "STUB: not implemented"
	return nil
}

func (g *abstractTaskGroup[T, E, O]) SubmitErr(tasks ...E) *abstractTaskGroup[T, E, O] {
	_ = "STUB: not implemented"
	return nil
}

func (g *abstractTaskGroup[T, E, O]) submit(task any) { _ = "STUB: not implemented"; return }

// Check if the context has been cancelled to prevent running tasks that are not needed

// Wrap the error with the context canceled error to reflect that the task was canceled.

// Invoke the task

type taskGroup struct {
	abstractTaskGroup[func(), func() error, struct{}]
}

func (g *taskGroup) Submit(tasks ...func()) TaskGroup {
	_ = "STUB: not implemented"
	return *new(TaskGroup)
}

func (g *taskGroup) SubmitErr(tasks ...func() error) TaskGroup {
	_ = "STUB: not implemented"
	return *new(TaskGroup)
}

func (g *taskGroup) Wait() error { _ = "STUB: not implemented"; return nil }

// This wait group could reach zero before the future is resolved if called in between tasks being submitted and the future being resolved.
// That's why we wait for the future to be resolved before waiting for the wait group.

type resultTaskGroup[O any] struct {
	abstractTaskGroup[func() O, func() (O, error), O]
}

func (g *resultTaskGroup[O]) Submit(tasks ...func() O) ResultTaskGroup[O] {
	_ = "STUB: not implemented"
	return nil
}

func (g *resultTaskGroup[O]) SubmitErr(tasks ...func() (O, error)) ResultTaskGroup[O] {
	_ = "STUB: not implemented"
	return nil
}

func (g *resultTaskGroup[O]) Wait() ([]O, error) { _ = "STUB: not implemented"; return nil, nil }

// This wait group could reach zero before the future is resolved if called in between tasks being submitted and the future being resolved.
// That's why we wait for the future to be resolved before waiting for the wait group.

func newTaskGroup(pool *pool, ctx context.Context) TaskGroup {
	_ = "STUB: not implemented"
	return *new(TaskGroup)
}

func newResultTaskGroup[O any](pool *pool, ctx context.Context) ResultTaskGroup[O] {
	_ = "STUB: not implemented"
	return nil
}
