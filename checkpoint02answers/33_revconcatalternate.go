// =============================================================
// EXERCISE: revconcatalternate
// QUESTION: Write RevConcatAlternate(slice1, slice2 []int) []int.
//           Interleave slices IN REVERSE ORDER, starting from the
//           LARGER slice (or slice1 if equal).
// =============================================================
//
// KEY INSIGHT — read the examples carefully:
//   RevConcatAlternate([1,2,3], [4,5,6])       → [3,6,2,5,1,4]
//   RevConcatAlternate([1,2,3], [4,5,6,7,8,9]) → [9,8,7,3,6,2,5,1,4]
//   RevConcatAlternate([1,2,3,9,8], [4,5])      → [8,9,3,2,5,1,4]
//   RevConcatAlternate([1,2,3], [])             → [3,2,1]
//
// PATTERN: Reverse both slices first, then apply ConcatAlternate logic
//   with the LONGER reversed slice leading.

package main

import "fmt"

// ─── THE SOLUTION ────────────────────────────────────────────

func RevConcatAlternate(slice1, slice2 []int) []int {

	// Step 1: Reverse both slices.
	rev1 := reverse(slice1)
	rev2 := reverse(slice2)

	// Step 2: Apply the same interleaving logic as ConcatAlternate,
	//         but with reversed slices.
	var result []int
	i, j := 0, 0
	s1Larger := len(rev1) >= len(rev2)

	for i < len(rev1) || j < len(rev2) {

		if !s1Larger {
			// rev2 is larger — it leads during the surplus.
			surplus := len(rev2) - len(rev1)
			if j < surplus {
				result = append(result, rev2[j])
				j++
			} else {
				if j < len(rev2) {
					result = append(result, rev2[j])
					j++
				}
				if i < len(rev1) {
					result = append(result, rev1[i])
					i++
				}
			}
		} else {
			// rev1 is larger or equal — it leads.
			if i < len(rev1) {
				result = append(result, rev1[i])
				i++
			}
			if j < len(rev2) {
				result = append(result, rev2[j])
				j++
			}
		}
	}

	return result
}

// reverse returns a new slice with elements in reverse order.
// It does NOT modify the original slice.
func reverse(s []int) []int {
	// DATA TYPE: []int — a new slice to hold the reversed version.
	result := make([]int, len(s))
	// "make([]int, len)" creates a slice of ints with the given length, all zeros.

	for i, v := range s {
		// Place v at position len-1-i (mirror position).
		// OPERATOR: - subtraction
		result[len(s)-1-i] = v
	}
	return result
}

// ─── DEMO ────────────────────────────────────────────────────
func main() {
	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	// → [3 6 2 5 1 4]

	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	// → [9 8 7 3 6 2 5 1 4]

	fmt.Println(RevConcatAlternate([]int{1, 2, 3, 9, 8}, []int{4, 5}))
	// → [8 9 3 2 5 1 4]

	fmt.Println(RevConcatAlternate([]int{1, 2, 3}, []int{}))
	// → [3 2 1]
}
