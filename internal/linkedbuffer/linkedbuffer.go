package linkedbuffer

import (
	"sync/atomic"
)

// LinkedBuffer implements an unbounded generic buffer that can be written to and read from concurrently.
// It is implemented using a linked list of buffers.
type LinkedBuffer[T any] struct {
	// Reader points to the buffer that is currently being read
	readBuffer *buffer[T]

	// Writer points to the buffer that is currently being written
	writeBuffer *buffer[T]

	maxCapacity int
	writeCount  atomic.Uint64
	readCount   atomic.Uint64
}

func NewLinkedBuffer[T any](initialCapacity, maxCapacity int) *LinkedBuffer[T] {
	_ = "STUB: not implemented"
	return nil
}

// Write writes values to the buffer
func (b *LinkedBuffer[T]) Write(value T) {
	_ = "STUB: not implemented"

	// Write elements
	return
}

// Increase next buffer capacity

// Retry writing

// Increment written count

// Read reads values from the buffer and returns the number of elements read
func (b *LinkedBuffer[T]) Read() (value T, err error) {
	_ = "STUB: not implemented"
	// Read element
	return *new(T), nil
}

// No more elements to read

// Move to next read buffer

// Retry reading

// Increment read count

// WriteCount returns the number of elements written to the buffer since it was created
func (b *LinkedBuffer[T]) WriteCount() uint64 { _ = "STUB: not implemented"; return 0 }

// ReadCount returns the number of elements read from the buffer since it was created
func (b *LinkedBuffer[T]) ReadCount() uint64 { _ = "STUB: not implemented"; return 0 }

// Len returns the number of elements in the buffer that haven't yet been read
func (b *LinkedBuffer[T]) Len() uint64 { _ = "STUB: not implemented"; return 0 }

// The writeCount counter wrapped around
