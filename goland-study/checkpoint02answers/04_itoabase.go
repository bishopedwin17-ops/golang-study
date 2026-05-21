// =============================================================
// EXERCISE: itoabase
// QUESTION: Write ItoaBase(nbr int, base int) string.
//           Convert an int to its string representation in
//           the given base (e.g. base 2 = binary, base 16 = hex).
// =============================================================

package main

import "fmt"

// ─── THE SOLUTION ────────────────────────────────────────────

// ItoaBase converts an int to a string in a given number base.
// nbr int  → the number to convert
// base int → the target base (2 to 16)
// string   → the result
func ItoaBase(nbr int, base int) string {

	// Guard: base must be between 2 and 16 inclusive.
	// An invalid base returns an empty string.
	// OPERATORS: < less-than, > greater-than, || logical OR
	if base < 2 || base > 16 {
		return ""
	}

	// "digits" is a string that acts as a lookup table.
	// Index 0 → '0', index 10 → 'a', index 15 → 'f', etc.
	// DATA TYPE: string — we'll index into it with [i] to get a byte.
	digits := "0123456789abcdef"

	// Handle the special case of zero.
	if nbr == 0 {
		return "0"
	}

	result := ""
	negative := false

	if nbr < 0 {
		negative = true
		nbr = -nbr // flip to positive
	}

	// Repeatedly divide by the base and collect remainders.
	for nbr != 0 {
		// nbr % base gives the current least-significant digit in that base.
		// e.g. 255 % 16 = 15  (which maps to 'f' in hex)
		remainder := nbr % base

		// digits[remainder] looks up the character for this digit.
		// e.g. digits[15] = 'f'
		// string(digits[remainder]) converts the single byte to a string.
		result = string(digits[remainder]) + result // prepend

		// Move to the next digit by dividing by the base.
		nbr = nbr / base
	}

	if negative {
		result = "-" + result
	}

	return result
}

// ─── DEMO ────────────────────────────────────────────────────
func main() {
	fmt.Println(ItoaBase(42, 2))   // → "101010"  (binary)
	fmt.Println(ItoaBase(255, 16)) // → "ff"       (hex)
	fmt.Println(ItoaBase(10, 8))   // → "12"       (octal)
	fmt.Println(ItoaBase(-5, 2))   // → "-101"
	fmt.Println(ItoaBase(0, 10))   // → "0"
	fmt.Println(ItoaBase(10, 1))   // → ""          (invalid base)
}
