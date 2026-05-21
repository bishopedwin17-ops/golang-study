// =============================================================
// EXERCISE: printrevcomb
// QUESTION: A PROGRAM that prints all unique combinations of
//           THREE different digits in DESCENDING order on one line.
//           First digit > second digit > third digit.
//           Combinations separated by ", ".
// =============================================================
//
// OUTPUT starts: 987, 986, 985, ..., 210
// Last combination is always 210 (smallest possible).
//
// NOTE: This is the REVERSE of printcomb.
//   printcomb:    012, 013, ..., 789   (ascending: a < b < c)
//   printrevcomb: 987, 986, ..., 210   (descending: a > b > c)

package main

import (
	"fmt"
	"os"
)

func main() {

	// Use os.Stdout for raw byte writing (piscine style).
	// "first" helps us avoid printing a comma before the very first combo.
	first := true

	// Three nested loops. The outer loop "a" is the LARGEST digit (9 down to 2).
	// The middle loop "b" is between a and c.
	// The inner loop "c" is the SMALLEST digit.
	//
	// CONSTRAINT: a > b > c, all different digits (0-9).
	//
	// We start a from 9 and go down to 2 (since b and c take at least 2 more values).
	for a := 9; a >= 2; a-- {
		// b must be less than a and greater than c (which is at least 0).
		// So b goes from a-1 down to 1.
		for b := a - 1; b >= 1; b-- {
			// c must be less than b.
			// So c goes from b-1 down to 0.
			for c := b - 1; c >= 0; c-- {

				// Print separator before every combo except the first.
				if !first {
					fmt.Fprintf(os.Stdout, ", ")
				}

				// Print the three digits as a 3-character number.
				// "%d%d%d" formats three integers without spaces.
				fmt.Fprintf(os.Stdout, "%d%d%d", a, b, c)
				first = false
			}
		}
	}

	// End with a newline.
	fmt.Println()
}

// ─── HOW TO RUN ──────────────────────────────────────────────
// go run 32_printrevcomb.go
// Output: 987, 986, 985, 984, 983, 982, 981, 980, 976, ..., 210
