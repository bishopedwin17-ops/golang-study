// =============================================================
// EXERCISE: fprime
// QUESTION: A PROGRAM that takes ONE positive integer and prints
//           its prime factors in ascending order, separated by "*".
//           Followed by \n.
//           If args != 1, invalid, or n <= 1 → print nothing.
// =============================================================
//
// EXAMPLE: fprime 225225
//   225225 = 3 × 3 × 5 × 5 × 7 × 11 × 13
//   Output: 3*3*5*5*7*11*13
//
// ALGORITHM: Trial division.
//   We try dividing n by 2, 3, 4, 5, ... until we've factored it fully.
//   When i divides n, we keep dividing (to catch repeated factors like 3*3).

package main

import (
	"fmt"
	"os"
)

// parsePositiveInt parses a string as a positive integer (> 1).
// Returns 0 on failure.
func parsePositiveInt(s string) int {
	if len(s) == 0 {
		return 0
	}
	n := 0
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		n = n*10 + int(s[i]-'0')
	}
	return n
}

func main() {

	if len(os.Args) != 2 {
		return // print nothing
	}

	n := parsePositiveInt(os.Args[1])

	// n must be > 1 to have prime factors.
	if n <= 1 {
		return
	}

	// ── Trial Division ────────────────────────────────────────
	result := ""    // will hold "3*3*5*..."
	first := true   // used to manage the "*" separator

	// Try every divisor "d" starting from 2.
	// We only need to go up to sqrt(n), PLUS handle the leftover.
	for d := 2; d*d <= n; d++ {
		// While d divides n evenly, extract it as a factor.
		for n%d == 0 {
			if !first {
				result += "*" // add separator before all factors except the first
			}
			// fmt.Sprintf formats an int as a string (like Itoa).
			result += fmt.Sprintf("%d", d)
			first = false

			// Divide n by d to remove this factor.
			// OPERATOR: / integer division
			n = n / d
		}
	}

	// If n is still > 1 after the loop, n itself is a prime factor.
	// e.g. if we started with 14 = 2 × 7, after dividing by 2 we get n=7.
	if n > 1 {
		if !first {
			result += "*"
		}
		result += fmt.Sprintf("%d", n)
	}

	fmt.Println(result)
}

// ─── HOW TO RUN ──────────────────────────────────────────────
// go run 15_fprime.go 225225   → 3*3*5*5*7*11*13
// go run 15_fprime.go 42       → 2*3*7
// go run 15_fprime.go 9539     → 9539  (itself is prime)
// go run 15_fprime.go 1        → (nothing)
// go run 15_fprime.go 0        → (nothing)
