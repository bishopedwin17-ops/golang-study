// =============================================================
// EXERCISE: cameltosnakecase
// QUESTION: Write CamelToSnakeCase(s string) string.
//           Convert camelCase to snake_case.
//           Return the string unchanged if it is not valid camelCase.
// Valid camelCase rules:
//   - Does not end on a capital letter  (e.g. "CamelCasE" → invalid)
//   - No two consecutive capitals       (e.g. "CamelCAse" → invalid)
//   - No digits or punctuation          (e.g. "camel1"    → invalid)
// =============================================================

package main

import "fmt"

// ─── THE SOLUTION ────────────────────────────────────────────

// CamelToSnakeCase converts a camelCase string to snake_case.
func CamelToSnakeCase(s string) string {

	// Convert the string to a slice of runes so we can index safely.
	// A rune is a Unicode code point — safely represents any character.
	// "[]rune(s)" is a type conversion: string → slice of runes.
	// DATA TYPE: []rune — a slice (dynamic array) of rune values.
	runes := []rune(s)

	// n is the number of characters (runes) in the string.
	n := len(runes)

	// Empty string → return it unchanged.
	if n == 0 {
		return s
	}

	// ── Validate the string ──────────────────────────────────
	// Walk through every rune and check the camelCase rules.
	for i, r := range runes {

		// Rule: no digits or punctuation anywhere.
		// A valid camelCase character must be a letter (a-z or A-Z).
		isLower := r >= 'a' && r <= 'z'
		isUpper := r >= 'A' && r <= 'Z'

		// OPERATOR: ! logical NOT (flips true to false and vice-versa)
		if !isLower && !isUpper {
			// Found a non-letter character → invalid camelCase.
			return s // return unchanged
		}

		// Rule: no two consecutive uppercase letters.
		// "i+1 < n" ensures we don't go out of bounds.
		// OPERATOR: + addition, < less-than
		if isUpper && i+1 < n && runes[i+1] >= 'A' && runes[i+1] <= 'Z' {
			return s // two capitals in a row → invalid
		}
	}

	// Rule: the last character must not be uppercase.
	if runes[n-1] >= 'A' && runes[n-1] <= 'Z' {
		return s // ends on capital → invalid
	}

	// ── Convert valid camelCase → snake_case ─────────────────
	// "result" builds the output string one rune at a time.
	result := ""

	for i, r := range runes {
		// If this is an uppercase letter (and not the very first character),
		// insert an underscore '_' before it.
		if r >= 'A' && r <= 'Z' && i != 0 {
			// string('_') converts the rune '_' to a one-character string.
			result += string('_')
		}
		// Append the current character (uppercase stays uppercase here).
		result += string(r)
	}

	return result
}

// ─── DEMO ────────────────────────────────────────────────────
func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))      // → Hello_World
	fmt.Println(CamelToSnakeCase("helloWorld"))      // → hello_World
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))// → camel_To_Snake_Case
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))// → unchanged (two caps)
	fmt.Println(CamelToSnakeCase("hey2"))            // → unchanged (digit)
	fmt.Println(CamelToSnakeCase("CamelCasE"))       // → unchanged (ends on cap)
}
