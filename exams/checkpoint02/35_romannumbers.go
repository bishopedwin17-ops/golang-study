// =============================================================
// EXERCISE: romannumbers
// QUESTION: A PROGRAM that converts a positive integer (1-3999)
//           to Roman numerals, printing both the calculation and result.
//           > 3999 or invalid → print "ERROR: cannot convert to roman digit"
// =============================================================
//
// ROMAN NUMERAL TABLE:
//   I=1, IV=4, V=5, IX=9, X=10, XL=40, L=50, XC=90,
//   C=100, CD=400, D=500, CM=900, M=1000
//
// ALGORITHM: Greedy — repeatedly subtract the largest roman value that fits.
//
// OUTPUT FORMAT for 123:
//   C+X+X+I+I+I
//   CXXIII
//
// For subtractive forms like 9 (IX), 4 (IV), the calculation uses (X-I) style.
//   OUTPUT for 999:
//   (M-C)+(C-X)+(X-I)
//   CMXCIX

package main

import (
	"fmt"
	"os"
)

// ─── LOOKUP TABLE ────────────────────────────────────────────
// We store pairs: (value, symbol, calcSymbol).
// For subtractive forms (CM, XC, etc.), calcSymbol shows "M-C" style.
// DATA TYPE: struct — groups related fields.
type romanEntry struct {
	value      int
	symbol     string // the actual roman numeral string
	calcSymbol string // how it appears in the calculation line
}

// The table in DESCENDING order (greedy needs largest first).
var romanTable = []romanEntry{
	{1000, "M", "M"},
	{900, "CM", "M-C"},
	{500, "D", "D"},
	{400, "CD", "D-C"},
	{100, "C", "C"},
	{90, "XC", "C-X"},
	{50, "L", "L"},
	{40, "XL", "L-X"},
	{10, "X", "X"},
	{9, "IX", "X-I"},
	{5, "V", "V"},
	{4, "IV", "V-I"},
	{1, "I", "I"},
}

// ─── HELPER: parsePositiveInt ────────────────────────────────
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

// ─── MAIN ────────────────────────────────────────────────────
func main() {

	if len(os.Args) != 2 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	n := parsePositiveInt(os.Args[1])

	// Valid range: 1 to 3999.
	if n < 1 || n > 3999 {
		fmt.Println("ERROR: cannot convert to roman digit")
		return
	}

	// ── Build calculation line and result string ──────────────
	calcParts := []string{} // e.g. ["M", "M", "M-C", ...]
	romanResult := ""

	remaining := n

	for _, entry := range romanTable {
		// How many times does this value fit?
		for remaining >= entry.value {
			// Does the calcSymbol need parentheses?
			// Subtractive forms (those with '-') need parens.
			needsParens := false
			for _, ch := range entry.calcSymbol {
				if ch == '-' {
					needsParens = true
					break
				}
			}

			if needsParens {
				calcParts = append(calcParts, "("+entry.calcSymbol+")")
			} else {
				calcParts = append(calcParts, entry.calcSymbol)
			}

			romanResult += entry.symbol
			remaining -= entry.value
		}
	}

	// Join calculation parts with "+".
	calcLine := ""
	for i, part := range calcParts {
		if i > 0 {
			calcLine += "+"
		}
		calcLine += part
	}

	fmt.Println(calcLine)
	fmt.Println(romanResult)
}

// ─── HOW TO RUN ──────────────────────────────────────────────
// go run 35_romannumbers.go 123   → C+X+X+I+I+I \n CXXIII
// go run 35_romannumbers.go 999   → (M-C)+(C-X)+(X-I) \n CMXCIX
// go run 35_romannumbers.go hello → ERROR: cannot convert to roman digit
