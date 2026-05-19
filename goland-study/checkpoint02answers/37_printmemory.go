// =============================================================
// EXERCISE: printmemory
// QUESTION: Write PrintMemory(arr [10]byte).
//           Print the memory in hex, 4 bytes per line.
//           Then print all ASCII graphic chars (printable),
//           replacing non-printable with '.'.
// =============================================================
//
// EXAMPLE: PrintMemory([10]byte{'h','e','l','l','o',16,21,'*'})
//   Hex output (4 bytes per line):
//     68 65 6c 6c
//     6f 10 15 2a
//     00 00
//   ASCII output:
//     hello..*..*
//   Wait — let me recheck the expected:
//     68 65 6c 6c   ← h(68) e(65) l(6c) l(6c)
//     6f 10 15 2a   ← o(6f) 16(10) 21(15) *(2a)
//     00 00         ← last two zero bytes
//     hello..*..*  ← non-printable → '.'
//
// PRINTABLE ASCII: characters with values 32 to 126 inclusive.

package main

import (
	"fmt"
	"os"
)

// ─── THE SOLUTION ────────────────────────────────────────────

// PrintMemory displays 10 bytes as hex rows + ASCII representation.
// arr [10]byte → a fixed-size ARRAY of 10 bytes.
// DATA TYPE: [10]byte — a fixed-size array (NOT a slice). Size is part of the type.
// DATA TYPE: byte — same as uint8, an unsigned 8-bit integer (0-255).
func PrintMemory(arr [10]byte) {

	// ── Print hex section: 4 bytes per line ─────────────────
	for i := 0; i < 10; i++ {

		// Print the hex value of arr[i].
		// "%02x" formats a byte as at least 2 hex digits, zero-padded.
		// e.g. byte 6 → "06", byte 104 → "68"
		// '0' = zero-pad, '2' = minimum 2 chars, 'x' = lowercase hex.
		fmt.Fprintf(os.Stdout, "%02x", arr[i])

		// After 4 bytes (or at the very last byte), print newline.
		// Otherwise, print a space between bytes.
		if (i+1)%4 == 0 || i == 9 {
			fmt.Fprintf(os.Stdout, "\n")
		} else {
			fmt.Fprintf(os.Stdout, " ")
		}
	}

	// ── Print ASCII section ──────────────────────────────────
	for i := 0; i < 10; i++ {
		b := arr[i]
		// ASCII graphic (printable) range: 32 (' ') to 126 ('~').
		// OPERATOR: >= greater-than-or-equal
		// OPERATOR: <= less-than-or-equal
		if b >= 32 && b <= 126 {
			// It is printable: output the character as-is.
			// string(b) converts the byte to a single-character string.
			fmt.Fprintf(os.Stdout, "%c", b)
		} else {
			// Non-printable: replace with dot '.'.
			fmt.Fprintf(os.Stdout, ".")
		}
	}
	fmt.Fprintf(os.Stdout, "\n")
}

// ─── DEMO ────────────────────────────────────────────────────
func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
	// Expected:
	//   68 65 6c 6c
	//   6f 10 15 2a
	//   00 00
	//   hello..*..*
}
