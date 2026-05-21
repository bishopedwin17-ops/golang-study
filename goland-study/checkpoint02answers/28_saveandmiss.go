// =============================================================
// EXERCISE: saveandmiss
// QUESTION: Write SaveAndMiss(arg string, num int) string.
//           Move through arg in sets of "num" characters.
//           SAVE the first set, SKIP the second, SAVE the third, etc.
//           If num <= 0 → return the original string unchanged.
// =============================================================
//
// EXAMPLE: SaveAndMiss("123456789", 3)
//   Group 1 (save):  "123"
//   Group 2 (skip):  "456"
//   Group 3 (save):  "789"
//   Result: "123789"
//
// EXAMPLE: SaveAndMiss("abcdefghijklmnopqrstuvwyz", 3)
//   Save: abc, skip: def, save: ghi, skip: jkl, save: mno, skip: pqr, save: stu, skip: vwy, save: z
//   Result: "abcghimnostuz"  (wait that's 13 chars — let's trust the spec)

package main

import "fmt"

// ─── THE SOLUTION ────────────────────────────────────────────

func SaveAndMiss(arg string, num int) string {

	// If num is 0 or negative → return the original string unchanged.
	if num <= 0 {
		return arg
	}

	result := ""

	// Convert string to rune slice for safe indexing.
	runes := []rune(arg)
	n := len(runes)

	// "save" tracks whether the current group should be saved (true) or skipped (false).
	// We start by saving.
	save := true

	// Walk through in chunks of "num".
	i := 0
	for i < n {
		// Determine end of this group.
		end := i + num
		if end > n {
			end = n
		}

		if save {
			// Save characters from index i to end (exclusive).
			for j := i; j < end; j++ {
				result += string(runes[j])
			}
		}
		// If !save → skip this group (do nothing).

		// Move to the next group.
		i = end
		// Toggle save/skip for the next group.
		// OPERATOR: ! logical NOT (flips true→false, false→true)
		save = !save
	}

	return result
}

// ─── DEMO ────────────────────────────────────────────────────
func main() {
	fmt.Println(SaveAndMiss("123456789", 3))                    // → 123789
	fmt.Println(SaveAndMiss("abcdefghijklmnopqrstuvwyz", 3))   // → abcghimnostuz
	fmt.Println(SaveAndMiss("", 3))                             // → (empty)
	fmt.Println(SaveAndMiss("hello you all ! ", 0))             // → hello you all !
	fmt.Println(SaveAndMiss("what is your name?", 0))          // → what is your name?
	fmt.Println(SaveAndMiss("go Exercise Save and Miss", -5))  // → go Exercise Save and Miss
}
