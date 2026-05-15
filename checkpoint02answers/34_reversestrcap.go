// =============================================================
// EXERCISE: reversestrcap
// QUESTION: A PROGRAM that takes ONE OR MORE arguments.
//           For each argument: uppercase the LAST letter of each
//           word, lowercase everything else.
//           Print result + \n per argument.
//           No arguments → print nothing.
// =============================================================
//
// EXAMPLE: "First SMALL TesT"
//   Words: [First, SMALL, TesT]
//   Apply: firsT, smalL, tesT
//   Output: "firsT smalL tesT"
//
// A "word" is a sequence of alphabetic/punctuation characters
// separated by spaces. Non-letter chars at the end of a word
// get the uppercase treatment only if they are letters.
// Actually, looking at the example: "IT'S" → "it'S"
// The LAST CHARACTER that is a letter gets uppercased.

package main

import (
	"fmt"
	"os"
)

// reverseCapArg processes one argument string.
func reverseCapArg(s string) string {

	runes := []rune(s)
	n := len(runes)

	// First, lowercase everything.
	for i := 0; i < n; i++ {
		runes[i] = toLower(runes[i])
	}

	// Then, find the last alphabetic character in each word and uppercase it.
	// A "word" boundary is a space character.
	// We scan left to right; when we hit a space or end, we find the last
	// letter in the previous word and uppercase it.

	// We'll track "inWord" and the index of the last letter seen.
	lastLetterIdx := -1
	inWord := false

	for i := 0; i <= n; i++ {

		atEnd := i == n
		isSpace := !atEnd && runes[i] == ' '

		if !atEnd && !isSpace {
			// We are inside a word.
			inWord = true
			// Check if this character is a letter.
			if isLetter(runes[i]) {
				lastLetterIdx = i // update last letter position in this word
			}
		} else {
			// Space or end of string: the word (if any) just ended.
			if inWord && lastLetterIdx >= 0 {
				// Uppercase the last letter of this word.
				runes[lastLetterIdx] = toUpper(runes[lastLetterIdx])
			}
			// Reset for the next word.
			inWord = false
			lastLetterIdx = -1
		}
	}

	return string(runes)
}

// isLetter returns true if ch is a-z or A-Z.
func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

// toLower converts a rune to lowercase if it is an uppercase letter.
func toLower(ch rune) rune {
	if ch >= 'A' && ch <= 'Z' {
		// Distance from 'A' to 'a' in ASCII is 32.
		// Adding 32 converts uppercase to lowercase.
		return ch + 32
	}
	return ch
}

// toUpper converts a rune to uppercase if it is a lowercase letter.
func toUpper(ch rune) rune {
	if ch >= 'a' && ch <= 'z' {
		return ch - 32
	}
	return ch
}

// ─── MAIN ────────────────────────────────────────────────────
func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		return // print nothing
	}

	for _, arg := range args {
		fmt.Println(reverseCapArg(arg))
	}
}

// ─── HOW TO RUN ──────────────────────────────────────────────
// go run 34_reversestrcap.go "First SMALL TesT"
// → firsT smalL tesT
