// =============================================================
// EXERCISE: findpairs
// QUESTION: A PROGRAM that finds all pairs of elements in an
//           integer array (given as "[1, 2, 3]") that sum to a
//           given target (given as "6").
//           Print: "Pairs with sum N: [[idx1 idx2] ...]"
//           Or appropriate error messages.
// =============================================================
//
// INPUT FORMAT: go run . "[1, 2, 3, 4, 5]" "6"
//   → Pairs with sum 6: [[0 4] [1 3]]
//
// ERROR MESSAGES:
//   "No pairs found."
//   "Invalid target sum."   (target has spaces or is not a single int)
//   "Invalid number: X"     (a number in the array is not an int)
//   "Invalid input."        (format doesn't match "[n, n, ...]" "target")

package main

import (
	"fmt"
	"os"
)

// ─── HELPER: parseIntStr ─────────────────────────────────────
// Parses a trimmed string as an integer. Returns (value, ok).
func parseIntStr(s string) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}
	i := 0
	sign := 1
	if s[i] == '-' {
		sign = -1
		i++
	} else if s[i] == '+' {
		i++
	}
	if i >= len(s) {
		return 0, false
	}
	n := 0
	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		n = n*10 + int(s[i]-'0')
	}
	return sign * n, true
}

// trimSpaces removes leading and trailing spaces from a string.
func trimSpaces(s string) string {
	start := 0
	end := len(s)
	for start < end && s[start] == ' ' {
		start++
	}
	for end > start && s[end-1] == ' ' {
		end--
	}
	return s[start:end]
}

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Invalid input.")
		return
	}

	arrStr := os.Args[1]  // e.g. "[1, 2, 3, 4, 5]"
	targetStr := os.Args[2] // e.g. "6"

	// ── Parse target ──────────────────────────────────────────
	// Target must be a single valid integer (no spaces within it).
	trimmedTarget := trimSpaces(targetStr)
	// If target has spaces → invalid target sum.
	for _, ch := range targetStr {
		if ch == ' ' {
			fmt.Println("Invalid target sum.")
			return
		}
	}
	target, ok := parseIntStr(trimmedTarget)
	if !ok {
		fmt.Println("Invalid target sum.")
		return
	}

	// ── Parse array string ───────────────────────────────────
	// Must start with '[' and end with ']'.
	if len(arrStr) < 2 || arrStr[0] != '[' || arrStr[len(arrStr)-1] != ']' {
		fmt.Println("Invalid input.")
		return
	}

	// Extract the inner content between [ and ].
	inner := arrStr[1 : len(arrStr)-1]

	// Split by comma.
	var nums []int
	// Walk through inner, splitting by ','.
	token := ""
	for i := 0; i <= len(inner); i++ {
		if i == len(inner) || inner[i] == ',' {
			t := trimSpaces(token)
			if t != "" {
				n, ok := parseIntStr(t)
				if !ok {
					fmt.Printf("Invalid number: %s\n", t)
					return
				}
				nums = append(nums, n)
			}
			token = ""
		} else {
			token += string(inner[i])
		}
	}

	// ── Find pairs ────────────────────────────────────────────
	// We need pairs (i, j) where i < j and nums[i] + nums[j] == target.
	// Each element used only once in a pair.
	type pair struct{ a, b int }
	var pairs []pair

	used := make([]bool, len(nums)) // tracks which indices are already in a pair

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			if used[j] {
				continue
			}
			if nums[i]+nums[j] == target {
				pairs = append(pairs, pair{i, j})
				used[i] = true
				used[j] = true
				break // each element used only once
			}
		}
	}

	if len(pairs) == 0 {
		fmt.Println("No pairs found.")
		return
	}

	// Print the pairs.
	fmt.Printf("Pairs with sum %d: [", target)
	for i, p := range pairs {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("[%d %d]", p.a, p.b)
	}
	fmt.Println("]")
}

// ─── HOW TO RUN ──────────────────────────────────────────────
// go run 38_findpairs.go "[1, 2, 3, 4, 5]" "6"
// → Pairs with sum 6: [[0 4] [1 3]]
//
// go run 38_findpairs.go "[1, 2, 3, 4, 5]" "10"
// → No pairs found.
