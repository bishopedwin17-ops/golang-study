// =============================================================
// EXERCISE: addprimesum
// QUESTION: A PROGRAM that takes ONE positive integer argument
//           and prints the sum of all prime numbers ≤ that integer.
//           If args != 1 or not a valid positive number → print "0\n".
// =============================================================
//
// EXAMPLE: addprimesum 5
//   Primes ≤ 5: 2, 3, 5 → sum = 10

package main

import (
	"fmt"
	"os"
)

// ─── HELPER: isPrime ─────────────────────────────────────────
// Returns true if n is a prime number.
// A prime is divisible only by 1 and itself, and must be > 1.
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	// Check divisors up to sqrt(n) using i*i <= n trick.
	for i := 2; i*i <= n; i++ {
		if n%i == 0 { // n is divisible by i → not prime
			return false
		}
	}
	return true
}

// ─── HELPER: parsePositiveInt ────────────────────────────────
// Converts a string to a positive int. Returns -1 if invalid.
func parsePositiveInt(s string) int {
	if len(s) == 0 {
		return -1
	}
	result := 0
	for i := 0; i < len(s); i++ {
		// Each character must be a digit.
		if s[i] < '0' || s[i] > '9' {
			return -1
		}
		result = result*10 + int(s[i]-'0')
	}
	// Must be > 0 (positive).
	if result <= 0 {
		return -1
	}
	return result
}

// ─── MAIN ────────────────────────────────────────────────────
func main() {

	// Validate argument count.
	if len(os.Args) != 2 {
		fmt.Println(0)
		return
	}

	// Parse the argument.
	n := parsePositiveInt(os.Args[1])
	if n == -1 {
		fmt.Println(0)
		return
	}

	// Sum all primes from 2 up to and including n.
	// DATA TYPE: int — holds the running sum.
	sum := 0

	// "i" starts at 2 (the smallest prime) and goes up to n.
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			// OPERATOR: += add and assign (sum = sum + i)
			sum += i
		}
	}

	fmt.Println(sum)
}

// ─── HOW TO RUN ──────────────────────────────────────────────
// go run 14_addprimesum.go 5   → 10   (2+3+5)
// go run 14_addprimesum.go 7   → 17   (2+3+5+7)
// go run 14_addprimesum.go -2  → 0
// go run 14_addprimesum.go     → 0
