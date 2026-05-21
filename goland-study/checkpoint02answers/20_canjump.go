// =============================================================
// EXERCISE: canjump
// QUESTION: Write CanJump(arr []uint) bool.
//           Given a slice where each value = exact steps forward,
//           can you reach the LAST index starting from index 0?
//           Return true if yes, false if no.
//           Single-element slice → true. Empty slice → false.
// =============================================================
//
// EXAMPLE: [2, 3, 1, 1, 4]
//   Start at index 0 → value is 2 → jump to index 2
//   At index 2 → value is 1 → jump to index 3
//   At index 3 → value is 1 → jump to index 4 (last!) → TRUE
//
// EXAMPLE: [3, 2, 1, 0, 4]
//   Start at 0 → jump 3 → index 3
//   At 3 → value is 0 → stuck! Can't move. → FALSE
//   (Note: we never even consider index 1 or 2 because we MUST
//    take exactly that many steps, not choose.)

package main

import "fmt"

// ─── THE SOLUTION ────────────────────────────────────────────

// CanJump simulates jumping through the array.
// arr []uint → slice of unsigned integers (uint cannot be negative).
// bool       → true if we can reach the end.
//
// DATA TYPE: uint — unsigned integer, only 0 and positive values.
// DATA TYPE: []uint — a slice of unsigned integers.
func CanJump(arr []uint) bool {

	// Empty slice → cannot even start.
	if len(arr) == 0 {
		return false
	}

	// Single element → we are already at the last index.
	if len(arr) == 1 {
		return true
	}

	// "pos" is our current position (index) in the array.
	// DATA TYPE: int — we use int (not uint) so we can compare with len().
	pos := 0

	// "last" is the index of the final element.
	// OPERATOR: - subtraction
	last := len(arr) - 1

	// Keep jumping until we reach the end or get stuck out of bounds.
	for pos != last {

		// How many steps must we take from here?
		// arr[pos] is the value at our current position.
		// DATA TYPE: uint → we convert to int with int() for arithmetic.
		steps := int(arr[pos])

		// Take exactly "steps" forward.
		pos = pos + steps

		// If we've overshot the array (gone past the last index) → false.
		// OPERATOR: > greater-than
		if pos > last {
			return false
		}

		// If steps == 0, we are stuck (infinite loop otherwise).
		if steps == 0 {
			return false
		}
	}

	return true
}

// ─── DEMO ────────────────────────────────────────────────────
func main() {
	fmt.Println(CanJump([]uint{2, 3, 1, 1, 4})) // → true
	fmt.Println(CanJump([]uint{3, 2, 1, 0, 4})) // → false (stuck at 0)
	fmt.Println(CanJump([]uint{0}))              // → true  (already at end)
	fmt.Println(CanJump([]uint{}))               // → false (empty)
}
