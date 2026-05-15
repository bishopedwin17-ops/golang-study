// =============================================================
// EXERCISE: expandstr
// QUESTION: A PROGRAM that takes ONE argument and prints it with
//           exactly THREE spaces between each word.
//           No leading/trailing whitespace. Followed by \n.
//           If args != 1 or no words → print nothing (no output at all).
// =============================================================

package main

import (
	"fmt"
	"os"
)

func main() {

	// We need exactly one argument (os.Args[0] is always the program name).
	if len(os.Args) != 2 {
		return // print nothing and exit
	}

	s := os.Args[1]

	// ── Extract words (non-whitespace sequences) ─────────────
	var words []string
	word := ""

	for _, ch := range s {
		// Any character that is NOT a space and NOT a tab is part of a word.
		if ch != ' ' && ch != '\t' {
			word += string(ch)
		} else {
			if len(word) > 0 {
				words = append(words, word)
				word = ""
			}
		}
	}
	if len(word) > 0 {
		words = append(words, word)
	}

	// If no words found → print nothing.
	if len(words) == 0 {
		return
	}

	// ── Join with THREE spaces ────────────────────────────────
	// "   " is a string literal containing exactly three space characters.
	result := ""
	for i, w := range words {
		result += w
		if i < len(words)-1 {
			result += "   " // three spaces between words
		}
	}

	fmt.Println(result)
}

// ─── HOW TO RUN ──────────────────────────────────────────────
// go run 10_expandstr.go "  only  it's harder   "
// Output: only   it's   harder
