// =============================================================
// EXERCISE: itoa
// QUESTION: Write a function Itoa(n int) string that converts
//           an integer like -42 into the string "-42".
// =============================================================

package main

import "fmt"

// ─── THE SOLUTION ────────────────────────────────────────────

// Itoa converts an int to its string representation.
// n int    → the number to convert, e.g. -42
// string   → the result, e.g. "-42"
func Itoa(n int) string {

	// Special case: 0 is just the string "0".
	// "==" is the equality operator.
	if n == 0 {
		return "0"
	}

	// "result" will hold the characters we build, in reverse order.
	// DATA TYPE: string — a sequence of characters (text).
	result := ""

	// "negative" remembers whether the original number was negative.
	// DATA TYPE: bool — can only be true or false.
	negative := false

	if n < 0 {
		negative = true
		// Flip n to positive so we can process its digits.
		// OPERATOR: * -1 multiplies by negative one → changes the sign.
		n = n * -1
		// Note: in Go we can also write: n = -n
	}

	// Extract digits from right to left using modulo and division.
	// OPERATOR: != means "not equal to"
	for n != 0 {
		// n % 10 gives the last digit.
		// e.g. 123 % 10 = 3
		// OPERATOR: % modulo (remainder of division)
		digit := n % 10

		// Convert the digit (int) to a rune character.
		// '0' has rune value 48. '0' + 3 = rune 51 = character '3'.
		// string(rune) converts a single rune to a one-character string.
		// OPERATOR: + adds two rune values together
		result = string(rune('0'+digit)) + result
		// We PREPEND (add to the front) because we extract digits right-to-left.

		// Remove the last digit by integer-dividing by 10.
		// e.g. 123 / 10 = 12  (integer division drops the remainder)
		// OPERATOR: / integer division
		n = n / 10
	}

	// If the number was negative, add the minus sign at the front.
	if negative {
		result = "-" + result
	}

	return result
}

// ─── DEMO ────────────────────────────────────────────────────
func main() {
	fmt.Println(Itoa(12345))    // → "12345"
	fmt.Println(Itoa(0))        // → "0"
	fmt.Println(Itoa(-1234))    // → "-1234"
	fmt.Println(Itoa(987654321))// → "987654321"
}
