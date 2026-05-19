// =============================================================
// EXERCISE: checknumber
// QUESTION: Write CheckNumber(arg string) bool.
//           Return true if the string contains at least one digit,
//           otherwise return false.
// =============================================================

package main

import "fmt"

// ─── THE SOLUTION ────────────────────────────────────────────

// CheckNumber checks if the string contains any digit character.
// arg string → the input text
// bool       → true if a digit is found, false otherwise
//
// DATA TYPE: bool — a boolean, either true or false.
func CheckNumber(arg string) bool {

	// Loop through every character (rune) in the string.
	// "_" discards the index (we only care about the value "ch").
	for _, ch := range arg {

		// Check if the character is a digit.
		// '0' (rune value 48) to '9' (rune value 57) are the digit runes.
		// OPERATOR: >= greater-than-or-equal
		// OPERATOR: <= less-than-or-equal
		// OPERATOR: && logical AND (both sides must be true)
		if ch >= '0' && ch <= '9' {
			// Found a digit — return true immediately.
			return true
		}
	}

	// No digit found after checking all characters.
	return false
}

// ─── DEMO ────────────────────────────────────────────────────
func main() {
	fmt.Println(CheckNumber("Hello"))  // → false
	fmt.Println(CheckNumber("Hello1")) // → true
	fmt.Println(CheckNumber(""))       // → false (empty has no digits)
	fmt.Println(CheckNumber("007"))    // → true
}
