// =============================================================
// EXERCISE: concatslice
// QUESTION: Write ConcatSlice(slice1, slice2 []int) []int.
//           Return a new slice that is slice1 followed by slice2.
// =============================================================

package main

import "fmt"

// ─── THE SOLUTION ────────────────────────────────────────────

// ConcatSlice joins two int slices into one.
// slice1 []int → first slice
// slice2 []int → second slice
// []int        → combined result
func ConcatSlice(slice1, slice2 []int) []int {

	// Method: use append to join them.
	// "append(slice1, slice2...)" appends ALL elements of slice2 onto slice1.
	// The "..." (called the ellipsis operator) UNPACKS a slice into individual
	// arguments. Without it, Go would try to append the whole slice as one element.
	//
	// DATA TYPE: []int — a slice (dynamic array) of ints.
	result := append(slice1, slice2...)

	return result
}

// ─── DEMO ────────────────────────────────────────────────────
func main() {
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{4, 5, 6}))       // → [1 2 3 4 5 6]
	fmt.Println(ConcatSlice([]int{}, []int{4, 5, 6, 7, 8, 9}))     // → [4 5 6 7 8 9]
	fmt.Println(ConcatSlice([]int{1, 2, 3}, []int{}))               // → [1 2 3]
}
