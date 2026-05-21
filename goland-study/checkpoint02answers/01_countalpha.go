// =============================================================
// EXERCISE: countalpha
// QUESTION: Write a function CountAlpha() that takes a string
//           and returns the number of alphabetic characters in it.
// =============================================================

// Every Go file starts with a "package" declaration.
// "package main" means this file is part of the main (runnable) program.
// If this were a library, we'd write "package piscine" instead.
package main

// "import" brings in other packages we need.
// "fmt" is the format package — it gives us Println, Printf, etc.
import "fmt"

// ─── THE SOLUTION ────────────────────────────────────────────

// CountAlpha is the function name. In Go, functions start with "func".
// (s string) → the parameter: "s" is the variable name, "string" is its type.
//              A string is a sequence of characters, e.g. "Hello".
// int         → the return type. "int" is a whole number (integer), e.g. 5.
func CountAlpha(s string) int {

	// "count" is a variable that will hold our running total.
	// ":=" means "declare AND assign at the same time".
	// "0" is the initial value (we start counting from zero).
	// DATA TYPE: int — whole numbers like 0, 1, -5, 42.
	count := 0

	// "for" starts a loop.
	// "range s" iterates over every character in the string "s".
	// "_" (blank identifier) discards the index (position) we don't need.
	// "ch" is the variable that holds each character in turn.
	// DATA TYPE: rune — a single Unicode character, e.g. 'H', 'e', '3'.
	//            A rune is really just an int32 under the hood.
	for _, ch := range s {

		// We check if the character is a letter:
		//   'a' <= ch && ch <= 'z'  → is it a lowercase letter?
		//   'A' <= ch && ch <= 'Z'  → is it an uppercase letter?
		//
		// OPERATORS USED:
		//   <=   less-than-or-equal comparison
		//   &&   logical AND (both conditions must be true)
		//   ||   logical OR  (at least one condition must be true)
		//
		// CHARACTERS (rune literals):
		//   'a'  the rune for the letter a (value 97 in ASCII)
		//   'z'  the rune for the letter z (value 122 in ASCII)
		//   'A'  the rune for the letter A (value 65 in ASCII)
		//   'Z'  the rune for the letter Z (value 90 in ASCII)
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {

			// "+=" means "add 1 to count and store the result back in count".
			// This is shorthand for: count = count + 1
			count++
			// "count++" is even shorter — it increments count by exactly 1.
		}
	}

	// "return" sends the final value back to whoever called this function.
	return count
}

// ─── DEMO main ───────────────────────────────────────────────
// "main" is the special entry-point function Go runs first.
func main() {
	// fmt.Println prints the value and adds a newline (\n) at the end.
	// Calling CountAlpha("Hello world") passes the string "Hello world" in.
	fmt.Println(CountAlpha("Hello world")) // → 10  (H,e,l,l,o,w,o,r,l,d)
	fmt.Println(CountAlpha("H e l l o"))  // → 5   (H,e,l,l,o — spaces ignored)
	fmt.Println(CountAlpha("H1e2l3l4o"))  // → 5   (digits don't count)
	fmt.Println(CountAlpha(""))           // → 0   (empty string has no letters)
}
