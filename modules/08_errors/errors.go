package errors_exercise

import (
	"errors"
	"fmt"
)

// ErrDivisionByZero is a sentinel error for division by zero.
// TODO: Define this using errors.New(...).
var ErrDivisionByZero = errors.New("division by zero")

// SafeDivide divides a by b. Returns ErrDivisionByZero if b == 0.
// TODO: Implement.
func SafeDivide(a, b float64) (float64, error) {
	// Your code here
	return 0, nil
}

// ParseAge parses a string to an int age.
// Returns an AgeError if the age is invalid (negative or > 150).
// Hint: use strconv.Atoi to convert string to int.
// TODO: Implement.
func ParseAge(s string) (int, error) {
	// Your code here
	return 0, nil
}

// AgeError is a custom error type.
// TODO: Implement the error interface for this struct.
type AgeError struct {
	Value   int
	Message string
}

func (e *AgeError) Error() string {
	// Your code here
	return ""
}

// WrapError wraps an error with context using fmt.Errorf and %w.
// e.g. WrapError("reading config", err) → "reading config: <err>"
// TODO: Implement.
func WrapError(context string, err error) error {
	return fmt.Errorf("%s: %w", context, err)
}
