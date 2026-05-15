// =============================================================
// EXERCISE: fifthandskip
// QUESTION: Write FifthAndSkip(str string) string.
//           Every 5 non-space characters form a group, separated
//           by spaces in output. The 6th character (after each group)
//           is SKIPPED/removed.
//           Spaces in input are ignored when counting.
//           If string < 5 non-space chars → return "Invalid Input\n".
//           If empty → return "\n".
// =============================================================
//
// EXAMPLE: "abcdefghijklmnopqrstuwxyz"
//   Non-space chars: a b c d e | f g h i j | k l m n o | p q r s t | u w x y z
//   Groups of 5:     [abcde]   [fghij]    [klmno]    [pqrst]    [uwxyz]
//   Output: "abcde ghijk mnopq stuwx z"
//   (The 6th character after each group is the one we skip.)

package main

import "fmt"

// ─── THE SOLUTION ────────────────────────────────────────────

func FifthAndSkip(str string) string {

	// Edge case: empty string.
	if len(str) == 0 {
		return "\n"
	}

	// First collect all non-space characters into a slice.
	// DATA TYPE: []rune — slice of runes, one per character.
	var chars []rune
	for _, ch := range str {
		// We skip spaces but keep everything else.
		if ch != ' ' {
			chars = append(chars, ch)
		}
	}

	// If fewer than 5 non-space characters → invalid.
	if len(chars) < 5 {
		return "Invalid Input\n"
	}

	// Now build the result:
	//   - Take 5 characters (a group).
	//   - Skip the 6th.
	//   - Separate groups with a space.
	result := ""
	groupCount := 0 // how many complete groups we've started

	i := 0
	for i < len(chars) {
		// Add a space between groups (but not before the first one).
		if groupCount > 0 {
			result += " "
		}

		// Take up to 5 characters for this group.
		groupSize := 0
		for groupSize < 5 && i < len(chars) {
			result += string(chars[i])
			i++
			groupSize++
		}

		groupCount++

		// Skip the 6th character (if it exists).
		if i < len(chars) {
			i++ // skip one character
		}
	}

	return result + "\n"
}

// ─── DEMO ────────────────────────────────────────────────────
func main() {
	fmt.Print(FifthAndSkip("abcdefghijklmnopqrstuwxyz"))
	// → abcde ghijk mnopq stuwx z

	fmt.Print(FifthAndSkip("This is a short sentence"))
	// → Thisi ashor sente ce

	fmt.Print(FifthAndSkip("1234"))
	// → Invalid Input

	fmt.Print(FifthAndSkip(""))
	// → (blank line)
}
