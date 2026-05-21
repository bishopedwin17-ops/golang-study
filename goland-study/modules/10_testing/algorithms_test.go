package testing_exercise

import "testing"

// --- Fibonacci tests (table-driven) ---

func TestFibonacci(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{10, 55},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Fibonacci(tt.n); got != tt.want {
				t.Errorf("Fibonacci(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(30)
	}
}

// --- IsPrime tests ---

func TestIsPrime(t *testing.T) {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23}
	notPrimes := []int{0, 1, 4, 6, 8, 9, 10, 25, 100}

	for _, n := range primes {
		t.Run("prime", func(t *testing.T) {
			if !IsPrime(n) {
				t.Errorf("IsPrime(%d) = false, want true", n)
			}
		})
	}
	for _, n := range notPrimes {
		t.Run("not_prime", func(t *testing.T) {
			if IsPrime(n) {
				t.Errorf("IsPrime(%d) = true, want false", n)
			}
		})
	}
}

// --- ReverseString tests ---

func TestReverseString(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"hello", "olleh"},
		{"", ""},
		{"a", "a"},
		{"racecar", "racecar"},
		{"Go!", "!oG"},
	}
	for _, tt := range tests {
		if got := ReverseString(tt.input); got != tt.want {
			t.Errorf("ReverseString(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}
