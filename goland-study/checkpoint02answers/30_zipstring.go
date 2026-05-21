// =============================================================
// EXERCISE: zipstring
// QUESTION: Write ZipString(s string) string.
//           Replace each sequence of identical characters with
//           the COUNT followed by the CHARACTER.
//           Only latin alphabet letters. Other chars stay as-is.
// =============================================================
//
// EXAMPLE: ZipString("YouuungFellllas")
//   Y → 1Y
//   o → 1o
//   u → 1u  (wait — "uuu" = 3u, but output shows "3u")
//   Actually let me re-read: "replace every character with the number
//   of duplicates and the character itself, deleting the extra duplications."
//   So "uuu" → "3u"  (count=3, char=u, collapse the 3 into 1 token)
//   "Fellll" → "1F 4l" = "1F1e4l"
//   Output for "YouuungFellllas": 1Y1o3u1n1g1F1e4l1a1s

package main

import "fmt"

// ─── THE SOLUTION ────────────────────────────────────────────

// ZipString compresses runs of repeated characters.
func ZipString(s string) string {

	if len(s) == 0 {
		return ""
	}

	runes := []rune(s) // convert to rune slice for safe iteration
	result := ""
	n := len(runes)
	i := 0

	for i < n {
		// Current character.
		current := runes[i]
		count := 1

		// Count how many times this character repeats consecutively.
		// "j" looks ahead.
		for i+count < n && runes[i+count] == current {
			count++
		}

		// Append "countCHAR" to result.
		// We convert count to its digit character(s).
		// For the piscine, counts are typically small (1-9), so we handle
		// multi-digit counts properly using our manual int→string conversion.
		result += intToStr(count) + string(current)

		// Jump past all the repeated characters.
		i += count
	}

	return result
}

// intToStr converts a positive integer to its string representation.
// (We avoid importing strconv to stay piscine-style.)
func intToStr(n int) string {
	if n == 0 {
		return "0"
	}
	result := ""
	for n > 0 {
		digit := n % 10
		result = string(rune('0'+digit)) + result
		n = n / 10
	}
	return result
}

// ─── DEMO ────────────────────────────────────────────────────
func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	// → 1Y1o3u1n1g1F1e4l1a1s

	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	// → 1T1h2e1 1q2u1i1c1k ...

	fmt.Println(ZipString("Helloo Therre!"))
	// → 1H1e2l2o1 1T1h1e2r1e1!
}
