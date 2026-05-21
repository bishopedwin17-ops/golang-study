// =============================================================
// EXERCISE: slice
// QUESTION: Write Slice(a []string, nbrs ...int) []string.
//           Return a sub-slice of "a" using the given indices.
//           One int → slice from that index to end.
//           Two ints → slice from first to second (exclusive).
//           Negative indices count from the END.
//           Invalid range → return nil.
// =============================================================
//
// EXAMPLES with a = ["coding","algorithm","ascii","package","golang"]:
//   Slice(a, 1)      → ["algorithm","ascii","package","golang"]
//   Slice(a, 2, 4)   → ["ascii","package"]
//   Slice(a, -3)     → ["ascii","package","golang"]   (-3 = index 2)
//   Slice(a, -2, -1) → ["package"]                    (-2=3, -1=4)
//   Slice(a, 2, 0)   → nil    (start > end → invalid)

package main

import "fmt"

// ─── THE SOLUTION ────────────────────────────────────────────

// Slice returns a sub-slice of "a" based on the given index/indices.
// "nbrs ...int" is a VARIADIC parameter: it accepts 0 or more ints.
// Inside the function, nbrs is a []int slice.
func Slice(a []string, nbrs ...int) []string {

	n := len(a)

	// Helper: convert a potentially negative index to a positive one.
	// e.g. -1 with len=5 → 5 + (-1) = 4
	normalize := func(idx int) int {
		if idx < 0 {
			return n + idx
		}
		return idx
	}

	if len(nbrs) == 1 {
		// One argument: slice from nbrs[0] to end.
		start := normalize(nbrs[0])
		if start < 0 || start >= n {
			return nil
		}
		return a[start:]

	} else if len(nbrs) == 2 {
		// Two arguments: slice from nbrs[0] to nbrs[1] (exclusive).
		start := normalize(nbrs[0])
		end := normalize(nbrs[1])

		// Invalid range.
		if start < 0 || end < 0 || start >= end || end > n {
			return nil
		}
		return a[start:end]
	}

	// 0 or >2 arguments → return nil.
	return nil
}

// ─── DEMO ────────────────────────────────────────────────────
func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}

	fmt.Printf("%#v\n", Slice(a, 1))     // → []string{"algorithm","ascii","package","golang"}
	fmt.Printf("%#v\n", Slice(a, 2, 4))  // → []string{"ascii","package"}
	fmt.Printf("%#v\n", Slice(a, -3))    // → []string{"ascii","package","golang"}
	fmt.Printf("%#v\n", Slice(a, -2, -1))// → []string{"package"}
	fmt.Printf("%#v\n", Slice(a, 2, 0))  // → []string(nil)
}
