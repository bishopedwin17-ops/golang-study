// =============================================================
// EXERCISE: cleanstr
// QUESTION: A PROGRAM (not a function) that takes ONE argument
//           and prints it with exactly one space between words,
//           no leading/trailing whitespace, followed by \n.
//           If args != 1 or no words → print only \n.
// =============================================================

package main

import (
	// "fmt" gives us Println and Print for output.
	"fmt"
	// "os" gives us os.Args — the command-line arguments.
	"os"
)

// ─── THE SOLUTION ────────────────────────────────────────────

func main() {

	// os.Args is a slice of strings: ["program_name", "arg1", "arg2", ...]
	// os.Args[0] is always the program name.
	// os.Args[1] would be the first user-provided argument.
	// len(os.Args) tells us how many arguments there are in total.
	//
	// We need EXACTLY 1 user argument → len(os.Args) must equal 2.
	// OPERATOR: != not equal
	if len(os.Args) != 2 {
		// Print just a newline and stop.
		fmt.Println()
		return // "return" exits the main function (ends the program).
	}

	// Grab the single argument.
	// DATA TYPE: string — text data.
	s := os.Args[1]

	// ── Extract words ─────────────────────────────────────────
	// We define a "word" as any sequence of non-whitespace characters.
	// A space ' ' (32) and tab '\t' (9) are both whitespace.

	// "words" collects all words we find.
	// DATA TYPE: []string — a slice (dynamic array) of strings.
	var words []string
	word := ""

	for _, ch := range s {
		// A tab '\t' or space ' ' is whitespace.
		isWhitespace := ch == ' ' || ch == '\t'

		if isWhitespace {
			if len(word) > 0 {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(ch)
		}
	}
	// Capture any trailing word.
	if len(word) > 0 {
		words = append(words, word)
	}

	// ── Print result ──────────────────────────────────────────
	if len(words) == 0 {
		// No words found in the argument → just a newline.
		fmt.Println()
		return
	}

	// Join words with exactly one space.
	result := ""
	for i, w := range words {
		result += w
		// Add space after every word EXCEPT the last.
		// OPERATOR: < less-than
		if i < len(words)-1 {
			result += " "
		}
	}

	// Print the result followed by a newline (Println adds \n automatically).
	fmt.Println(result)
}

// ─── HOW TO RUN ──────────────────────────────────────────────
// go run 09_cleanstr.go "  only    it's  harder   "
// Output: only it's harder
//
// go run 09_cleanstr.go ""
// Output: (blank line)
//
// go run 09_cleanstr.go
// Output: (blank line)   ← wrong number of args
