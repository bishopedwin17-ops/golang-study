// =============================================================
// EXERCISE: notdecimal
// QUESTION: Write NotDecimal(dec string) string.
//           Takes a float string like "174.2" and returns it as
//           an integer string with the decimal point removed
//           (multiplied by 10^n where n = number of decimal digits).
//           - No decimal or only ".0" → return the number + \n.
//           - Empty → return "\n".
//           - Not a valid number → return it unchanged + \n.
// =============================================================
//
// EXAMPLES:
//   "0.1"       → "1\n"         (0.1 × 10 = 1)
//   "174.2"     → "1742\n"      (174.2 × 10 = 1742)
//   "0.1255"    → "1255\n"      (0.1255 × 10000 = 1255)
//   "-19.525856"→ "-19525856\n"
//   "1952"      → "1952\n"      (no decimal → return as-is)
//   "-0.0f00d00"→ "-0.0f00d00\n" (not a valid number)

package main

import "fmt"

// ─── THE SOLUTION ────────────────────────────────────────────

func NotDecimal(dec string) string {

	// Empty → just newline.
	if len(dec) == 0 {
		return "\n"
	}

	// ── Validate: check if it looks like a valid decimal number ──
	// A valid number: optional leading '-', digits, optional '.' then digits.
	runes := []rune(dec)
	n := len(runes)
	i := 0

	// Optional sign.
	negative := false
	if runes[i] == '-' {
		negative = true
		i++
	} else if runes[i] == '+' {
		i++
	}

	// Collect integer part (digits before the dot).
	intPart := ""
	for i < n && runes[i] >= '0' && runes[i] <= '9' {
		intPart += string(runes[i])
		i++
	}

	// If we have no digits yet, invalid.
	if len(intPart) == 0 {
		return dec + "\n"
	}

	// Check for a decimal point.
	fracPart := ""
	if i < n && runes[i] == '.' {
		i++ // skip the dot
		// Collect fractional digits.
		for i < n && runes[i] >= '0' && runes[i] <= '9' {
			fracPart += string(runes[i])
			i++
		}
		// Any remaining characters → invalid.
		if i < n {
			return dec + "\n"
		}
	} else if i < n {
		// Unexpected character after the integer part → invalid.
		return dec + "\n"
	}

	// ── Build the result ──────────────────────────────────────

	// If no fractional part, or fractional part is all zeros → return as-is.
	allZeroFrac := true
	for _, ch := range fracPart {
		if ch != '0' {
			allZeroFrac = false
			break
		}
	}

	if len(fracPart) == 0 || allZeroFrac {
		// Just the integer part (plus sign if negative).
		prefix := ""
		if negative {
			prefix = "-"
		}
		return prefix + intPart + "\n"
	}

	// Combine integer and fractional parts, dropping trailing zeros from frac.
	// e.g. "1.20525856" → intPart="1", fracPart="20525856"
	// Result: "120525856"
	// But first strip trailing zeros from fracPart.
	trimFrac := fracPart
	for len(trimFrac) > 0 && trimFrac[len(trimFrac)-1] == '0' {
		trimFrac = trimFrac[:len(trimFrac)-1]
	}

	prefix := ""
	if negative {
		prefix = "-"
	}

	// The integer result = intPart + fracPart (no dot, no leading zeros in intPart).
	// If intPart == "0", drop it (e.g. "0.1" → "1", not "01").
	combined := intPart + trimFrac
	// Remove leading zeros from combined (keeping at least one digit).
	for len(combined) > 1 && combined[0] == '0' {
		combined = combined[1:]
	}

	return prefix + combined + "\n"
}

// ─── DEMO ────────────────────────────────────────────────────
func main() {
	fmt.Print(NotDecimal("0.1"))         // → 1
	fmt.Print(NotDecimal("174.2"))       // → 1742
	fmt.Print(NotDecimal("0.1255"))      // → 1255
	fmt.Print(NotDecimal("1.20525856"))  // → 120525856
	fmt.Print(NotDecimal("-0.0f00d00"))  // → -0.0f00d00  (invalid)
	fmt.Print(NotDecimal(""))            // → (blank line)
	fmt.Print(NotDecimal("-19.525856"))  // → -19525856
	fmt.Print(NotDecimal("1952"))        // → 1952
}
