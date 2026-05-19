// =============================================================
// EXERCISE: union
// QUESTION: A PROGRAM that takes TWO arguments and prints,
//           without duplicates, all characters that appear in
//           EITHER string, in the order they appear on the
//           command line (arg1 first, then arg2).
//           Followed by \n.
//           If args != 2 → print just \n.
// =============================================================
//
// EXAMPLE:
//   arg1 = "zpadinton"
//   arg2 = "paqefwtdjetyiytjneytjoeyjnejeyj"
//   Union (no duplicates, arg1 order then arg2 new chars):
//   → "zpadintoqefwjy"

package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println() // just a newline
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	// "seen" maps each character to whether we have already output it.
	// DATA TYPE: map[rune]bool
	seen := make(map[rune]bool)

	result := ""

	// Process s1 first, then s2.
	// By concatenating them (s1 + s2) we can use a single loop.
	// DATA TYPE: string — s1 + s2 joins two strings together.
	// OPERATOR: + string concatenation
	combined := s1 + s2

	for _, ch := range combined {
		// If we haven't seen this character yet, add it to the result.
		if !seen[ch] {
			result += string(ch)
			seen[ch] = true // mark as seen
		}
	}

	fmt.Println(result)
}

// ─── HOW TO RUN ──────────────────────────────────────────────
// go run 13_union.go zpadinton paqefwtdjetyiytjneytjoeyjnejeyj
// Output: zpadintoqefwjy
