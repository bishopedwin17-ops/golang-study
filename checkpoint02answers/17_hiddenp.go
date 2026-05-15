// =============================================================
// EXERCISE: hiddenp
// QUESTION: A PROGRAM that takes TWO arguments (s1 and s2).
//           Check if s1 is "hidden" in s2: every character of s1
//           appears in s2 in the same order (not necessarily consecutive).
//           Print "1\n" if hidden, "0\n" if not.
//           Empty s1 → always "1".
//           If args != 2 → print nothing.
// =============================================================
//
// DIFFERENCE from wdmatch: hiddenp always prints 1 or 0.
// wdmatch prints the word or nothing.

package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		return // print nothing
	}

	s1 := os.Args[1] // the string to find
	s2 := os.Args[2] // the string to search in

	// An empty s1 is always hidden in any s2.
	if len(s1) == 0 {
		fmt.Println(1)
		return
	}

	// "j" is our current position in s2.
	j := 0

	// Walk through each character of s1 and try to match it in s2.
	for _, ch := range s1 {

		// Advance j until we find the character ch in s2.
		for j < len(s2) && rune(s2[j]) != ch {
			j++
		}

		// If we ran out of s2 without finding ch → not hidden.
		if j >= len(s2) {
			fmt.Println(0) // print 0 with newline
			return
		}

		// Matched: move j forward past this character.
		j++
	}

	// All characters of s1 were matched in order → hidden!
	fmt.Println(1)
}

// ─── HOW TO RUN ──────────────────────────────────────────────
// go run 17_hiddenp.go "fgex.;" "tyf34gdgf;'ektufjhgdgex.;.;rtjynur6"  → 1
// go run 17_hiddenp.go "abc" "btarc"   → 0
// go run 17_hiddenp.go "DD" "DABC"     → 0  (only one D in s2)
// go run 17_hiddenp.go "" "anything"   → 1  (empty s1)
