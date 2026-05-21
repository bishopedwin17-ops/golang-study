// =============================================================
// EXERCISE: grouping
// QUESTION: A PROGRAM that takes TWO arguments:
//           1. A pattern like "(a)" or "(e|n)"
//           2. A sentence to search in
//           Print each word containing the pattern, numbered.
//           The "|" operator means "or" (match either side).
//           If args != 2, pattern invalid, sentence empty, or
//           no matches → print nothing.
// =============================================================
//
// EXAMPLE: grouping "(a)" "I'm heavy, jumpsuit is on steady"
//   Words containing 'a': heavy, steady, heavy → numbered 1,2,3
//
// EXAMPLE: grouping "(e|n)" "I currently have 4 windows opened"
//   Words containing 'e' OR 'n': currently(e), currently(n), have(e), ...

package main

import (
	"fmt"
	"os"
)

// ─── HELPER: containsSubstr ──────────────────────────────────
// Returns true if "word" contains "sub" as a substring.
func containsSubstr(word, sub string) bool {
	if len(sub) == 0 {
		return true
	}
	if len(sub) > len(word) {
		return false
	}
	wRunes := []rune(word)
	sRunes := []rune(sub)
	for i := 0; i <= len(wRunes)-len(sRunes); i++ {
		match := true
		for j := 0; j < len(sRunes); j++ {
			if wRunes[i+j] != sRunes[j] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}

// ─── HELPER: extractWords ─────────────────────────────────────
// Splits a string into words separated by spaces.
func extractWords(s string) []string {
	var words []string
	word := ""
	for _, ch := range s {
		if ch == ' ' {
			if len(word) > 0 {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(ch)
		}
	}
	if len(word) > 0 {
		words = append(words, word)
	}
	return words
}

func main() {

	if len(os.Args) != 3 {
		return // print nothing
	}

	pattern := os.Args[1] // e.g. "(a)" or "(e|n)"
	sentence := os.Args[2]

	// Validate pattern: must be "(something)" with optional "|" inside.
	// The pattern must start with '(' and end with ')'.
	if len(pattern) < 3 || pattern[0] != '(' || pattern[len(pattern)-1] != ')' {
		return // invalid pattern
	}

	// Extract the inner content between ( and ).
	inner := pattern[1 : len(pattern)-1] // e.g. "a" or "e|n"

	if len(inner) == 0 {
		return
	}

	// Split inner by '|' to get search terms.
	// DATA TYPE: []string
	var terms []string
	term := ""
	for _, ch := range inner {
		if ch == '|' {
			if len(term) == 0 {
				return // invalid: empty term
			}
			terms = append(terms, term)
			term = ""
		} else {
			term += string(ch)
		}
	}
	if len(term) == 0 {
		return // invalid
	}
	terms = append(terms, term)

	// Empty sentence → print nothing.
	if len(sentence) == 0 {
		return
	}

	// Extract words from the sentence.
	words := extractWords(sentence)

	// For each word, check if it matches ANY of the terms.
	// Each match prints a separate numbered line.
	count := 0
	for _, w := range words {
		for _, t := range terms {
			if containsSubstr(w, t) {
				count++
				fmt.Printf("%d: %s\n", count, w)
			}
		}
	}
}

// ─── HOW TO RUN ──────────────────────────────────────────────
// go run 39_grouping.go "(a)" "I'm heavy, jumpsuit is on steady, Lighter when I'm lower, higher when I'm heavy"
// → 1: heavy
//   2: steady
//   3: heavy
