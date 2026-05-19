// =============================================================
// EXERCISE: chunk
// QUESTION: Write Chunk(slice []int, size int).
//           Split "slice" into sub-slices of length "size".
//           Print each sub-slice group.
//           If size == 0 → print only \n.
//           If slice is empty → print [].
// =============================================================
//
// EXAMPLE: Chunk([0,1,2,3,4,5,6,7], 3)
//   Groups: [0,1,2], [3,4,5], [6,7]
//   Output: [[0 1 2] [3 4 5] [6 7]]

package main

import "fmt"

// ─── THE SOLUTION ────────────────────────────────────────────

// Chunk splits a slice into sub-slices of the given size and prints them.
// slice []int → the input integers
// size int    → the chunk size
func Chunk(slice []int, size int) {

	// size == 0 is invalid → print just a newline.
	if size == 0 {
		fmt.Println()
		return
	}

	// Empty slice → print [].
	if len(slice) == 0 {
		fmt.Println("[]")
		return
	}

	// "chunks" will hold all sub-slices.
	// DATA TYPE: [][]int — a slice of slices of ints (a 2D slice).
	var chunks [][]int

	// Walk through "slice" in steps of "size".
	// "i" starts at 0 and jumps by "size" each iteration.
	for i := 0; i < len(slice); i += size {

		// "end" is the index one past the last element of this chunk.
		end := i + size

		// Make sure "end" doesn't go past the slice length.
		// OPERATOR: > greater-than
		if end > len(slice) {
			end = len(slice)
		}

		// slice[i:end] creates a sub-slice from index i up to (not including) end.
		// DATA TYPE: []int — a slice expression returns a slice.
		// OPERATOR: : the slice operator, as in slice[start:end]
		chunk := slice[i:end]

		// Append this chunk to our result.
		chunks = append(chunks, chunk)
	}

	// fmt.Println knows how to print slices and slices-of-slices nicely.
	fmt.Println(chunks)
}

// ─── DEMO ────────────────────────────────────────────────────
func main() {
	Chunk([]int{}, 10)                    // → []
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0) // → (blank line)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3) // → [[0 1 2] [3 4 5] [6 7]]
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5) // → [[0 1 2 3 4] [5 6 7]]
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4) // → [[0 1 2 3] [4 5 6 7]]
}
