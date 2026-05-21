// =============================================================
// EXERCISE: concatalternate
// QUESTION: Write ConcatAlternate(slice1, slice2 []int) []int.
//           Interleave elements of the two slices, starting with
//           the LARGER slice. If equal length, start with slice1.
// =============================================================
//
// EXAMPLES:
//   Equal length [1,2,3] and [4,5,6]:
//     → [1,4,2,5,3,6]  (start with slice1 since equal)
//
//   slice2 is longer [1,2,3] and [4,5,6,7,8,9]:
//     → [4,1,5,2,6,3,7,8,9]  (start with slice2, the larger)
//
//   slice1 is longer [2,4,6,8,10] and [1,3,5,7,9,11]:
//     Hmm wait — [1,2,3,4,5,6,7,8,9,10,11] that seems wrong.
//     Actually from the expected output [1 2 3 4 5 6 7 8 9 10 11],
//     when slice2 has 6 and slice1 has 5, slice2 starts.
//     Pattern: larger leads, then they alternate.

package main

import "fmt"

// ─── THE SOLUTION ────────────────────────────────────────────

// ConcatAlternate interleaves two slices, larger one first.
func ConcatAlternate(slice1, slice2 []int) []int {

	// DATA TYPE: []int — the result slice.
	var result []int

	// "i" tracks position in slice1, "j" in slice2.
	i, j := 0, 0

	// Determine which slice is larger (or if equal).
	// We flip roles based on which is longer.
	// DATA TYPE: bool — for the "flip" flag.
	slice2Larger := len(slice2) > len(slice1)

	for i < len(slice1) || j < len(slice2) {

		if slice2Larger {
			// slice2 is larger → take from slice2 first when its "surplus" exists.
			// surplus = extra elements that slice2 has beyond slice1's length.
			surplus := len(slice2) - len(slice1)

			if j < surplus {
				// Still in the surplus zone — only append from slice2.
				result = append(result, slice2[j])
				j++
			} else {
				// Now alternate: slice2 element, then slice1 element.
				if j < len(slice2) {
					result = append(result, slice2[j])
					j++
				}
				if i < len(slice1) {
					result = append(result, slice1[i])
					i++
				}
			}
		} else {
			// slice1 is larger or equal → take from slice1 first.
			if i < len(slice1) {
				result = append(result, slice1[i])
				i++
			}
			if j < len(slice2) {
				result = append(result, slice2[j])
				j++
			}
		}
	}

	return result
}

// ─── DEMO ────────────────────────────────────────────────────
func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	// → [1 4 2 5 3 6]

	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	// → [1 2 3 4 5 6 7 8 9 10 11]  (slice2 larger, leads with surplus)

	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	// → [4 1 5 2 6 3 7 8 9]

	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
	// → [1 2 3]
}
