// =============================================================
// EXERCISE: findprevprime
// QUESTION: Write FindPrevPrime(nb int) int.
//           Return the largest prime number ≤ nb.
//           Return 0 if no prime exists ≤ nb.
// =============================================================

package main

import "fmt"

// ─── HELPER: isPrime ─────────────────────────────────────────

// isPrime checks whether n is a prime number.
// A prime is a number > 1 that has no divisors other than 1 and itself.
func isPrime(n int) bool {

	// Numbers ≤ 1 are not prime by definition.
	if n <= 1 {
		return false
	}

	// Check divisors from 2 up to the square root of n.
	// We only need to go up to sqrt(n) because if n has a factor larger
	// than sqrt(n), the paired factor would be smaller than sqrt(n) and
	// we'd have already caught it.
	//
	// OPERATOR: * multiplication
	// Instead of calling math.Sqrt (to avoid imports), we use i*i <= n.
	for i := 2; i*i <= n; i++ {
		// n % i == 0 means i divides n evenly → n is NOT prime.
		// OPERATOR: % modulo (remainder)
		// OPERATOR: == equality check
		if n%i == 0 {
			return false
		}
	}

	return true
}

// ─── THE SOLUTION ────────────────────────────────────────────

// FindPrevPrime finds the largest prime ≤ nb.
// nb int → the upper limit
// int    → the result prime, or 0 if none
func FindPrevPrime(nb int) int {

	// Start from nb and count DOWN until we find a prime.
	// OPERATOR: >= greater-than-or-equal
	for nb >= 2 {
		if isPrime(nb) {
			return nb // found it!
		}
		// Decrement nb by 1 to try the next smaller number.
		// "nb--" is shorthand for nb = nb - 1
		nb--
	}

	// No prime found at or below the given number.
	return 0
}

// ─── DEMO ────────────────────────────────────────────────────
func main() {
	fmt.Println(FindPrevPrime(5))  // → 5  (5 itself is prime)
	fmt.Println(FindPrevPrime(4))  // → 3  (4 is not prime; 3 is)
	fmt.Println(FindPrevPrime(10)) // → 7
	fmt.Println(FindPrevPrime(1))  // → 0  (no prime ≤ 1)
}
