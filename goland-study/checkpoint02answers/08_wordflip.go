// =============================================================
// EXERCISE: wordflip
// QUESTION: Write WordFlip(str string) string.
//           Return the words of the string in REVERSE order.
//           Trim spaces, collapse multiple spaces between words.
//           If empty or only spaces → return "Invalid Output\n".
//           Each result must end with a newline \n.
// =============================================================

package main

import "fmt"

// ─── THE SOLUTION ────────────────────────────────────────────

// WordFlip reverses the order of words in a string.
func WordFlip(str string) string {

	// Step 1: Split the string into individual words.
	//         We do this manually to avoid importing "strings".

	// "words" is a slice of strings — a dynamic list of strings.
	// DATA TYPE: []string — a slice (resizable array) of string values.
	// "nil" means the slice starts empty (no elements, no memory allocated yet).
	var words []string

	// "word" accumulates the characters of the current word being read.
	word := ""

	// Walk through every character in str.
	// "range str" gives index (i) and character (ch) for each iteration.
	// DATA TYPE: rune — a single Unicode character.
	for _, ch := range str {

		// A space character has rune value 32.
		// We use it as a word separator.
		if ch == ' ' {
			// If we have accumulated some characters, this word is done.
			// len(word) > 0 checks that the word is not empty.
			if len(word) > 0 {
				// "append" adds an element to the END of a slice.
				// It returns the new (possibly larger) slice.
				words = append(words, word)
				// Reset "word" to empty for the next word.
				word = ""
			}
			// If ch is a space and "word" is empty, we just skip (multiple spaces).
		} else {
			// Not a space: add this character to the current word.
			// "+=" appends the string representation of ch to word.
			word += string(ch)
		}
	}
	// After the loop, if there is a leftover word (string didn't end in space):
	if len(word) > 0 {
		words = append(words, word)
	}

	// Step 2: If no words were found → invalid.
	// len(words) returns the number of elements in the slice.
	if len(words) == 0 {
		return "Invalid Output\n"
	}

	// Step 3: Build the result by joining words in REVERSE order.
	result := ""

	// "i" starts at the LAST index (len-1) and counts DOWN to 0.
	// OPERATOR: - subtraction
	// OPERATOR: >= greater-than-or-equal
	for i := len(words) - 1; i >= 0; i-- {
		result += words[i]
		// Add a space between words, but NOT after the last one.
		// OPERATOR: != not-equal
		if i != 0 {
			result += " "
		}
	}

	// Each result ends with a newline character.
	// "\n" is an escape sequence meaning "newline".
	return result + "\n"
}

// ─── DEMO ────────────────────────────────────────────────────
func main() {
	fmt.Print(WordFlip("First second last"))   // → "last second First\n"
	fmt.Print(WordFlip(""))                    // → "Invalid Output\n"
	fmt.Print(WordFlip("     "))               // → "\n"  (only spaces → no words)
	fmt.Print(WordFlip(" hello  all  of  you!"))// → "you! of all hello\n"
}
