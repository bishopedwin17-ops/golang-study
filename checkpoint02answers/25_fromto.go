// =============================================================
// EXERCISE: fromto
// QUESTION: Write FromTo(from int, to int) string.
//           Return a string showing numbers from "from" to "to",
//           separated by ", ". Each number zero-padded to 2 digits.
//           Appended with \n.
//           If either number is < 0 or > 99 → return "Invalid\n".
// =============================================================
//
// EXAMPLES:
//   FromTo(1, 10)  → "01, 02, 03, 04, 05, 06, 07, 08, 09, 10\n"
//   FromTo(10, 1)  → "10, 09, 08, 07, 06, 05, 04, 03, 02, 01\n"
//   FromTo(10, 10) → "10\n"
//   FromTo(100,10) → "Invalid\n"

package main

import "fmt"

// ─── HELPER: zeroPad ─────────────────────────────────────────
// Returns a 2-digit string for a number 0-99.
// e.g. zeroPad(5) → "05"
//      zeroPad(42) → "42"
func zeroPad(n int) string {
	if n < 10 {
		// Single digit: prepend a "0".
		// string(rune('0'+n)) converts the digit to a character.
		return "0" + string(rune('0'+n))
	}
	// Two digits: convert manually using integer division.
	tens := n / 10   // e.g. 42 / 10 = 4
	ones := n % 10   // e.g. 42 % 10 = 2
	return string(rune('0'+tens)) + string(rune('0'+ones))
}

// ─── THE SOLUTION ────────────────────────────────────────────

func FromTo(from int, to int) string {

	// Validate: both numbers must be in [0, 99].
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid\n"
	}

	result := ""
	first := true // used to control the ", " separator

	if from <= to {
		// Count UP from from to to.
		for i := from; i <= to; i++ {
			if !first {
				result += ", "
			}
			result += zeroPad(i)
			first = false
		}
	} else {
		// Count DOWN from from to to.
		for i := from; i >= to; i-- {
			if !first {
				result += ", "
			}
			result += zeroPad(i)
			first = false
		}
	}

	return result + "\n"
}

// ─── DEMO ────────────────────────────────────────────────────
func main() {
	fmt.Print(FromTo(1, 10))   // → 01, 02, ..., 10
	fmt.Print(FromTo(10, 1))   // → 10, 09, ..., 01
	fmt.Print(FromTo(10, 10))  // → 10
	fmt.Print(FromTo(100, 10)) // → Invalid
}
