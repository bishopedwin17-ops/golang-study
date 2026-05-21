// =============================================================
// EXERCISE: inter
// QUESTION: A PROGRAM that takes TWO arguments and prints,
//           without duplicates, all characters that appear in
//           BOTH strings, in the order they appear in the FIRST.
//           Followed by \n.
//           If args != 2 → print nothing.
// =============================================================
//
// EXAMPLE:
//   arg1 = "padinton"
//   arg2 = "paqefwtdjetyiytjneytjoeyjnejeyj"
//   Characters in arg1 that also appear in arg2 (no duplicates):
//   p, a, d, i, n, t, o  → "padinto"

package main

import (
	"fmt"
	"os"
)

func main() {

	// We need exactly 2 arguments (so len(os.Args) == 3, because index 0 is
	// the program name itself).
	if len(os.Args) != 3 {
		return // print nothing
	}

	s1 := os.Args[1] // first string
	s2 := os.Args[2] // second string

	// ── Build a "seen in s2" lookup ───────────────────────────
	// We use a map: a data structure that maps keys to values.
	// DATA TYPE: map[rune]bool
	//   - key:   rune (a character)
	//   - value: bool (true = this character is in s2)
	//
	// "make" allocates and initialises the map.
	inS2 := make(map[rune]bool)

	for _, ch := range s2 {
		// Setting a map entry: inS2[ch] = true
		// This records that character ch appears in s2.
		inS2[ch] = true
	}

	// ── Walk s1 and collect characters that are also in s2 ───
	// "printed" tracks characters we have already output, so we
	// don't print duplicates.
	printed := make(map[rune]bool)

	result := ""
	for _, ch := range s1 {
		// inS2[ch] looks up whether ch is in the map (returns true/false).
		// !printed[ch] means "we haven't printed this character yet".
		// OPERATOR: ! logical NOT
		if inS2[ch] && !printed[ch] {
			result += string(ch)
			// Mark this character as already printed.
			printed[ch] = true
		}
	}

	// fmt.Println prints result and automatically appends \n.
	fmt.Println(result)
}

// ─── HOW TO RUN ──────────────────────────────────────────────
// go run 12_inter.go "padinton" "paqefwtdjetyiytjneytjoeyjnejeyj"
// Output: padinto
