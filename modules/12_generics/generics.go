package generics

// Number is a type constraint for numeric types.
type Number interface {
	int | int32 | int64 | float32 | float64
}

// Sum returns the total of all elements in a slice.
// TODO: Implement as a generic function using the Number constraint.
func Sum[T Number](nums []T) T {
	// Your code here
	var zero T
	return zero
}

// Contains returns true if item is in slice.
// TODO: Implement as a generic function using the comparable constraint.
func Contains[T comparable](slice []T, item T) bool {
	// Your code here
	return false
}

// Map applies a transform function to each element and returns a new slice.
// TODO: Implement. Signature: func Map[T, U any](slice []T, f func(T) U) []U
func Map[T, U any](slice []T, f func(T) U) []U {
	// Your code here
	return nil
}

// Stack is a generic LIFO stack.
// TODO: Implement Push, Pop, and Len methods.
type Stack[T any] struct {
	items []T
}

// Push adds an item to the top.
func (s *Stack[T]) Push(item T) {
	// Your code here
}

// Pop removes and returns the top item. Returns false if empty.
func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	// Your code here
	return zero, false
}

// Len returns the number of items in the stack.
func (s *Stack[T]) Len() int {
	// Your code here
	return 0
}
