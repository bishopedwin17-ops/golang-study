// =============================================================
// EXERCISE: atoi
// QUESTION: Write a function Atoi(s string) int that converts
//           a string like "123" into the integer 123.
//           Handle + and - signs. Return 0 for invalid input.
// =============================================================

package main

import "fmt"

// ─── THE SOLUTION ────────────────────────────────────────────

// Atoi converts a string to an int, mimicking strconv.Atoi.
// s string → the input, e.g. "-42" or "hello"
// int      → the result, e.g. -42 or 0
func Atoi(s string) int {

	// "result" will accumulate the number we are building.
	// "sign" tracks whether the number is positive (+1) or negative (-1).
	// DATA TYPE: int — whole numbers (positive, negative, or zero).
	result := 0
	sign := 1 // assume positive by default

	// "i" is the index (position) we start reading from.
	// We begin at 0 (the very first character of the string).
	i := 0

	// len(s) returns the number of bytes in s.
	// If the string is empty there is nothing to convert → return 0.
	if len(s) == 0 {
		return 0
	}

	// Check for a sign character at position 0.
	// s[0]  → accesses the first byte of the string (type: byte / uint8).
	// byte literals like '+' and '-' are single-quoted.
	if s[0] == '+' {
		// The number is explicitly positive. Skip past the '+'.
		// "i++" increments i by 1.
		i++
	} else if s[0] == '-' {
		// The number is negative. Flip sign to -1.
		sign = -1
		i++ // skip past the '-'
	}

	// After handling the sign, there must be at least one digit.
	// If i has already reached the end of the string, it is invalid.
	if i >= len(s) {
		return 0
	}

	// Now walk through each remaining character and build the number.
	// "for i < len(s)" → keep looping while i is less than the string length.
	// OPERATOR: <  means "less than"
	for i < len(s) {

		// s[i] gives the byte (character) at position i.
		// To check if it is a digit, we compare it to '0' and '9'.
		// ASCII values: '0'=48, '1'=49, … '9'=57.
		if s[i] < '0' || s[i] > '9' {
			// Non-digit found → the string is invalid → return 0.
			// OPERATOR: ||  logical OR
			return 0
		}

		// Convert the ASCII character to its numeric value.
		// e.g. '5' has ASCII value 53; 53 - 48 ('0') = 5.
		// "int(s[i] - '0')" → first subtracts bytes (gives a byte 0-9),
		//                      then converts that byte to int.
		digit := int(s[i] - '0')

		// Shift the result left by one decimal place and add the new digit.
		// e.g. if result=12 and digit=3 → result = 12*10 + 3 = 123
		// OPERATOR: *  multiplication
		// OPERATOR: +  addition
		result = result*10 + digit

		// Move to the next character.
		i++
	}

	// Apply the sign and return.
	// OPERATOR: *  multiplication (sign is either 1 or -1)
	return sign * result
}

// ─── DEMO ────────────────────────────────────────────────────
func main() {
	fmt.Println(Atoi("12345"))   // → 12345
	fmt.Println(Atoi("+1234"))   // → 1234
	fmt.Println(Atoi("-1234"))   // → -1234
	fmt.Println(Atoi("012 345")) // → 0  (space is not a digit)
	fmt.Println(Atoi("Hello!"))  // → 0  (letters are not digits)
	fmt.Println(Atoi("++1234"))  // → 0  (double sign → invalid)
	fmt.Println(Atoi("--1234"))  // → 0
	fmt.Println(Atoi(""))        // → 0  (empty)
}
