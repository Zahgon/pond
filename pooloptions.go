package pond

import (
	"context"
)

type Option func(*pool)

// WithContext sets the context for the pool.
func WithContext(ctx context.Context) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithQueueSize sets the max number of elements that can be queued in the pool.
func WithQueueSize(size int) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithNonBlocking sets the pool to be non-blocking when the queue is full.
// This option is only effective when the queue size is set.
func WithNonBlocking(nonBlocking bool) Option { _ = "STUB: not implemented"; return *new(Option) }

// WithoutPanicRecovery disables panic interception inside worker goroutines.
// When this option is enabled, panics inside tasks will propagate just like regular goroutines.
func WithoutPanicRecovery() Option { _ = "STUB: not implemented"; return *new(Option) }
