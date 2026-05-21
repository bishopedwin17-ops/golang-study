// =============================================================
// EXERCISE: wdmatch
// QUESTION: A PROGRAM that takes TWO arguments.
//           Check if every character of s1 can be found, IN ORDER,
//           in s2 (not necessarily consecutive).
//           If yes → print s1 + \n.
//           If no or wrong args → print nothing.
// =============================================================
//
// EXAMPLE:
//   s1 = "faya"
//   s2 = "fgvvfdxcacpolhyghbreda"
//   Can we find f, a, y, a in order in s2?
//   f✓ a✓ y✓ a✓  → YES → print "faya"
//
//   s1 = "error"
//   s2 = "rrerrrfiiljdfxjyuifrrvcoojh"
//   e? → not found → NO → print nothing

package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		return // print nothing
	}

	s1 := os.Args[1] // the pattern we want to find
	s2 := os.Args[2] // the string we search through

	// "j" is our position in s2.
	// We advance j every time we match a character from s1.
	// DATA TYPE: int — integer index
	j := 0

	// Walk through every character in s1.
	for _, ch := range s1 {

		// Advance j through s2 until we find a match for ch.
		// OPERATOR: && logical AND
		// len(s2) returns the byte length of s2.
		for j < len(s2) && rune(s2[j]) != ch {
			j++ // move to next position in s2
		}

		// If j reached the end of s2 without finding ch → no match.
		if j >= len(s2) {
			return // print nothing and exit
		}

		// We found ch at position j. Move past it for the next search.
		j++
	}

	// All characters of s1 were found in order in s2.
	fmt.Println(s1)
}

// ─── HOW TO RUN ──────────────────────────────────────────────
// go run 16_wdmatch.go faya fgvvfdxcacpolhyghbreda  → faya
// go run 16_wdmatch.go faya fgvvfdxcacpolhyghbred   → (nothing)
// go run 16_wdmatch.go error rrerrrfiiljdfxjyuifrrvcoojh → (nothing)
