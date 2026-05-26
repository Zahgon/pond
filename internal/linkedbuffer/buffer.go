package linkedbuffer

import (
	"errors"
)

var ErrEOF = errors.New("EOF")

// buffer implements a generic buffer that can store any type of data.
// It is not thread-safe and should be used with a mutex.
// It is used by LinkedBuffer to store data and is not intended to be used directly.
type buffer[T any] struct {
	data           []T
	nextWriteIndex int
	nextReadIndex  int
	next           *buffer[T]
	zero           T
}

func newBuffer[T any](capacity int) *buffer[T] { _ = "STUB: not implemented"; return nil }

// Cap returns the capacity of the buffer.
func (b *buffer[T]) Cap() int {
	_ = "STUB: not implemented"

	// Len returns the number of elements in the buffer.
	return 0
}

func (b *buffer[T]) Len() int { _ = "STUB: not implemented"; return 0 }

// Write writes a value to the buffer.
// If the buffer is full, it returns an EOF error.
func (b *buffer[T]) Write(value T) error { _ = "STUB: not implemented"; return nil }

// Buffer is full

// Read reads a value from the buffer.
// If the buffer is empty, it returns an EOF error.
// If the buffer has been read completely, it returns an EOF error.
func (b *buffer[T]) Read() (value T, err error) { _ = "STUB: not implemented"; return *new(T), nil }

// Buffer read completely, return EOF error

// Remove reference to read value to prevent memory leaks caused by
// holding references to submitted tasks.
// See https://github.com/alitto/pond/issues/110
