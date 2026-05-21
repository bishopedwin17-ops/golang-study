// =============================================================
// EXERCISE: thirdtimeisacharm
// QUESTION: Write ThirdTimeIsACharm(str string) string.
//           Return every THIRD character of the string.
//           Counting starts at 1 (so chars at position 3,6,9,...).
//           End with \n.
//           Empty string or no third character → return "\n".
// =============================================================
//
// EXAMPLE: "123456789"
//   Position: 1=1, 2=2, 3=3, 4=4, 5=5, 6=6, 7=7, 8=8, 9=9
//   Every 3rd:      3            6            9
//   Output: "369\n"
//
// EXAMPLE: "a b c d e f"   (with spaces counted as characters!)
//   Position: 1=a 2=' ' 3=b 4=' ' 5=c 6=' ' 7=d 8=' ' 9=e 10=' ' 11=f
//   Every 3rd:      b             ' '           e
//   Output: "b e\n"

package main

import "fmt"

// ─── THE SOLUTION ────────────────────────────────────────────

func ThirdTimeIsACharm(str string) string {

	if len(str) == 0 {
		return "\n"
	}

	result := ""
	// "pos" is a 1-based counter.
	// DATA TYPE: int — integer counter.
	pos := 1

	for _, ch := range str {
		// Every time pos is a multiple of 3, collect this character.
		// OPERATOR: % modulo (remainder of division)
		// pos%3 == 0 means pos is divisible by 3.
		if pos%3 == 0 {
			result += string(ch)
		}
		// Increment position counter.
		pos++
	}

	// If no third character was found, result is empty → return just "\n".
	return result + "\n"
}

// ─── DEMO ────────────────────────────────────────────────────
func main() {
	fmt.Print(ThirdTimeIsACharm("123456789")) // → "369\n"
	fmt.Print(ThirdTimeIsACharm(""))          // → "\n"
	fmt.Print(ThirdTimeIsACharm("a b c d e f")) // → "b e\n"  (spaces counted!)
	fmt.Print(ThirdTimeIsACharm("12"))        // → "\n"  (no 3rd char)
}
