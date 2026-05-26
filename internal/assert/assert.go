package assert

import "testing"

/**
 * Asserts that the expected and actual values are equal.
 */
func Equal(t *testing.T, expected interface{}, actual interface{}) {
	_ = "STUB: not implemented"
	return
}

/**
 * Asserts that the actual value is true.
 */
func True(t *testing.T, actual bool) { _ = "STUB: not implemented"; return }

/**
 * Asserts that the function panics with the expected object.
 */
func PanicsWith(t *testing.T, expected any, f func()) { _ = "STUB: not implemented"; return }

/**
 * Asserts that the function panics with the expected error.
 */
func PanicsWithError(t *testing.T, expected string, f func()) { _ = "STUB: not implemented"; return }

/**
 * Asserts that the function not panics.
 */
func NotPanics(t *testing.T, f func()) { _ = "STUB: not implemented"; return }
