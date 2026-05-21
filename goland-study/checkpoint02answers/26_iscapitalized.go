// =============================================================
// EXERCISE: iscapitalized
// QUESTION: Write IsCapitalized(s string) bool.
//           Return true if EVERY word starts with an uppercase
//           letter OR a non-alphabetic character.
//           Return false if any word starts with a lowercase letter.
//           Return false if the string is empty.
// =============================================================
//
// EXAMPLES:
//   "Hello! How are you?"  → false  ("are" and "you" start lowercase)
//   "Hello How Are You"    → true
//   "Whats 4this 100K?"    → true   ("4this" starts with digit, OK)
//   "!!!!Whatsthis4"       → true   ("!!!!" is non-alpha, then "Whatsthis4" starts W)

package main

import "fmt"

// ─── THE SOLUTION ────────────────────────────────────────────

func IsCapitalized(s string) bool {

	if len(s) == 0 {
		return false
	}

	// "newWord" is true when the NEXT character will be the start of a word.
	// At the beginning of the string, we are at the start of a word.
	newWord := true

	for _, ch := range s {

		if ch == ' ' {
			// A space means the next non-space character starts a new word.
			newWord = true
		} else if newWord {
			// We are at the start of a word.
			// Check if this character is a lowercase letter.
			isLower := ch >= 'a' && ch <= 'z'
			if isLower {
				// Word starts with lowercase → not capitalized.
				return false
			}
			// Upper case or non-alpha at word start → OK.
			// The word has started, so the next character is no longer "newWord".
			newWord = false
		}
		// If !newWord, we are in the middle of a word → don't check.
	}

	return true
}

// ─── DEMO ────────────────────────────────────────────────────
func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))  // → false
	fmt.Println(IsCapitalized("Hello How Are You"))    // → true
	fmt.Println(IsCapitalized("Whats 4this 100K?"))    // → true
	fmt.Println(IsCapitalized("Whatsthis4"))           // → true  (one word, starts W)
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))       // → true
	fmt.Println(IsCapitalized(""))                     // → false
}
