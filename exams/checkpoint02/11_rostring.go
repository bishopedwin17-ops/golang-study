// =============================================================
// EXERCISE: rostring
// QUESTION: A PROGRAM that takes ONE argument and rotates it
//           one word to the LEFT (first word moves to the end).
//           Words are separated by one space in the output.
//           A "word" = sequence of ALPHANUMERIC characters.
//           Non-alphanumeric chars are treated as separators but
//           are KEPT in output at their relative positions.
//           If args != 1 → print just \n.
// =============================================================
//
// HOW ROTATION WORKS:
//   Input:  "Let there be light"
//   Words:  [Let, there, be, light]
//   Rotate: [there, be, light, Let]
//   Output: "there be light Let"
//
//   Input:  "     AkjhZ zLKIJz , 23y"
//   Alphanumeric words: [AkjhZ, zLKIJz, 23y]
//   Non-alpha separators: [, ] stays
//   Output: "zLKIJz , 23y AkjhZ"

package main

import (
	"fmt"
	"os"
)

// isAlnum checks whether a rune is a letter (a-z, A-Z) or digit (0-9).
// DATA TYPE: rune — a single Unicode character.
// DATA TYPE: bool — true or false.
func isAlnum(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') ||
		(ch >= 'A' && ch <= 'Z') ||
		(ch >= '0' && ch <= '9')
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println() // print just a newline
		return
	}

	s := os.Args[1]

	// Convert to slice of runes for safe character-by-character access.
	runes := []rune(s)
	n := len(runes)

	// ── Collect "tokens" preserving order ────────────────────
	// A token is either a WORD (alphanumeric) or a NON-WORD (spaces/punctuation).
	// We store each token as a string.
	// DATA TYPE: []string — slice of strings
	type token struct {
		text  string // the actual characters
		isWord bool   // true if this token is an alphanumeric word
	}

	// "tokens" is a slice of token structs.
	// DATA TYPE: struct — a custom type grouping related fields together.
	var tokens []token
	i := 0

	for i < n {
		if isAlnum(runes[i]) {
			// Collect all consecutive alphanumeric characters → a word.
			word := ""
			for i < n && isAlnum(runes[i]) {
				word += string(runes[i])
				i++
			}
			tokens = append(tokens, token{word, true})
		} else {
			// Collect all consecutive non-alphanumeric characters → separator.
			sep := ""
			for i < n && !isAlnum(runes[i]) {
				sep += string(runes[i])
				i++
			}
			tokens = append(tokens, token{sep, false})
		}
	}

	// ── Find all word tokens ──────────────────────────────────
	// We need to rotate: move the first word to the end.
	var wordIndices []int // positions of word tokens inside "tokens" slice
	for idx, t := range tokens {
		if t.isWord {
			// "append" adds the index to our list.
			wordIndices = append(wordIndices, idx)
		}
	}

	// If there are no words or only one word, no rotation needed visually,
	// but we still rebuild the output.
	if len(wordIndices) > 1 {
		// Take the first word's text and move it to the last word's slot.
		firstWordText := tokens[wordIndices[0]].text
		// Shift all word texts left by one position.
		for i := 0; i < len(wordIndices)-1; i++ {
			tokens[wordIndices[i]].text = tokens[wordIndices[i+1]].text
		}
		// Put the first word at the end.
		tokens[wordIndices[len(wordIndices)-1]].text = firstWordText
	}

	// ── Rebuild the output string ─────────────────────────────
	result := ""
	for _, t := range tokens {
		result += t.text
	}

	// Trim leading and trailing spaces/newlines for clean output.
	// We do this manually by finding the first and last non-space rune.
	trimmed := ""
	rResult := []rune(result)
	start := 0
	end := len(rResult) - 1

	// Move start forward past leading spaces.
	for start <= end && rResult[start] == ' ' {
		start++
	}
	// Move end backward past trailing spaces.
	for end >= start && rResult[end] == ' ' {
		end--
	}
	for _, ch := range rResult[start : end+1] {
		trimmed += string(ch)
	}

	fmt.Println(trimmed)
}

// ─── HOW TO RUN ──────────────────────────────────────────────
// go run 11_rostring.go "Let there be light"
// Output: there be light Let
