package functions

import "fmt"

// Divide returns (result, error). If b == 0, return an error.
// Hint: use fmt.Errorf("...") to create an error.
// TODO: Implement.
func Divide(a, b float64) (float64, error) {
	// Your code here
	return 0, nil
}

// SumAll accepts a variadic number of ints and returns the total.
// TODO: Implement.
func SumAll(nums ...int) int {
	// Your code here
	return 0
}

// MakeAdder returns a closure that adds `n` to whatever is passed to it.
// e.g. add5 := MakeAdder(5); add5(3) == 8
// TODO: Implement.
func MakeAdder(n int) func(int) int {
	// Your code here
	return nil
}

// MinMax returns the minimum and maximum values in nums.
// Use named return values: (min, max int)
// TODO: Implement.
func MinMax(nums []int) (min, max int) {
	// Your code here
	return
}

// keep this — used internally
var _ = fmt.Errorf
