// =============================================================
// EXERCISE: weareunique
// QUESTION: Write WeAreUnique(str1, str2 string) int.
//           Count characters NOT shared between both strings
//           (no repeating characters in the count).
//           Both empty → return -1.
//           No unique chars → return 0.
// =============================================================
//
// EXAMPLE: WeAreUnique("foo", "boo")
//   Unique chars in "foo":  f, o  (deduplicated)
//   Unique chars in "boo":  b, o  (deduplicated)
//   Characters NOT in BOTH: f (only in str1), b (only in str2)
//   Count = 2 → return 2
//
// EXAMPLE: WeAreUnique("abc", "def")
//   str1 unique: a,b,c   str2 unique: d,e,f
//   Nothing shared → all 6 are unique → return 6

package main

import "fmt"

// ─── THE SOLUTION ────────────────────────────────────────────

func WeAreUnique(str1, str2 string) int {

	// Both empty → special case.
	if len(str1) == 0 && len(str2) == 0 {
		return -1
	}

	// Build a set of unique characters in str1.
	// DATA TYPE: map[rune]bool — maps a character to true if present.
	set1 := make(map[rune]bool)
	for _, ch := range str1 {
		set1[ch] = true
	}

	// Build a set of unique characters in str2.
	set2 := make(map[rune]bool)
	for _, ch := range str2 {
		set2[ch] = true
	}

	// Count characters in set1 that are NOT in set2.
	count := 0
	for ch := range set1 {
		// map lookup: set2[ch] returns false if ch is not in set2.
		if !set2[ch] {
			count++
		}
	}

	// Count characters in set2 that are NOT in set1.
	for ch := range set2 {
		if !set1[ch] {
			count++
		}
	}

	return count
}

// ─── DEMO ────────────────────────────────────────────────────
func main() {
	fmt.Println(WeAreUnique("foo", "boo")) // → 2  (f and b are unique)
	fmt.Println(WeAreUnique("", ""))       // → -1
	fmt.Println(WeAreUnique("abc", "def")) // → 6  (nothing shared)
}
