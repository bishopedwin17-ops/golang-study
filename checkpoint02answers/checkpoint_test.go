// =============================================================
// TEST FILE — Top 10 Most Common Checkpoint 02 Exercises
//
// HOW TO RUN: go test ./checkpoint02answers/ -v
//
// Each test is "table-driven": we define a table of input/output
// pairs and loop through them. This is idiomatic Go testing.
//
// DATA TYPE: testing.T — the Go test runner object.
//   t.Errorf() reports a test failure with a message.
//   t.Logf()   prints a message (only shown on failure or -v flag).
// =============================================================

package main

import "testing"

// ─── 1. ATOI ─────────────────────────────────────────────────

func TestAtoi(t *testing.T) {
	// "cases" is a slice of anonymous structs.
	// Each struct has an "input" (the string) and "want" (expected int).
	cases := []struct {
		input string
		want  int
	}{
		{"12345", 12345},
		{"+1234", 1234},
		{"-1234", -1234},
		{"0", 0},
		{"", 0},
		{"Hello!", 0},
		{"++1234", 0},
		{"--5", 0},
		{"012 345", 0}, // space in middle → invalid
	}

	for _, c := range cases {
		// Call our Atoi function.
		got := Atoi(c.input)
		// OPERATOR: != not equal
		if got != c.want {
			// t.Errorf marks the test as failed and prints a message.
			// %q prints a string with quotes around it (e.g. "hello").
			t.Errorf("Atoi(%q) = %d, want %d", c.input, got, c.want)
		}
	}
}

// ─── 2. ITOA ─────────────────────────────────────────────────

func TestItoa(t *testing.T) {
	cases := []struct {
		input int
		want  string
	}{
		{12345, "12345"},
		{0, "0"},
		{-1234, "-1234"},
		{987654321, "987654321"},
		{-1, "-1"},
	}

	for _, c := range cases {
		got := Itoa(c.input)
		if got != c.want {
			t.Errorf("Itoa(%d) = %q, want %q", c.input, got, c.want)
		}
	}
}

// ─── 3. COUNTALPHA ───────────────────────────────────────────

func TestCountAlpha(t *testing.T) {
	cases := []struct {
		input string
		want  int
	}{
		{"Hello world", 10},
		{"H e l l o", 5},
		{"H1e2l3l4o", 5},
		{"", 0},
		{"12345", 0},
		{"abc", 3},
	}

	for _, c := range cases {
		got := CountAlpha(c.input)
		if got != c.want {
			t.Errorf("CountAlpha(%q) = %d, want %d", c.input, got, c.want)
		}
	}
}

// ─── 4. CHECKNUMBER ──────────────────────────────────────────

func TestCheckNumber(t *testing.T) {
	cases := []struct {
		input string
		want  bool
	}{
		{"Hello", false},
		{"Hello1", true},
		{"", false},
		{"007", true},
		{"!@#$", false},
		{"abc3def", true},
	}

	for _, c := range cases {
		got := CheckNumber(c.input)
		if got != c.want {
			t.Errorf("CheckNumber(%q) = %v, want %v", c.input, got, c.want)
		}
	}
}

// ─── 5. FINDPREVPRIME ────────────────────────────────────────

func TestFindPrevPrime(t *testing.T) {
	cases := []struct {
		input int
		want  int
	}{
		{5, 5},   // 5 is prime
		{4, 3},   // 4 is not, 3 is
		{10, 7},
		{2, 2},   // smallest prime
		{1, 0},   // no prime ≤ 1
		{0, 0},
		{-5, 0},
		{20, 19},
	}

	for _, c := range cases {
		got := FindPrevPrime(c.input)
		if got != c.want {
			t.Errorf("FindPrevPrime(%d) = %d, want %d", c.input, got, c.want)
		}
	}
}

// ─── 6. WORDFLIP ─────────────────────────────────────────────

func TestWordFlip(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{"First second last", "last second First\n"},
		{"", "Invalid Output\n"},
		{"     ", "Invalid Output\n"},
		{" hello  all  of  you!", "you! of all hello\n"},
		{"one", "one\n"},
	}

	for _, c := range cases {
		got := WordFlip(c.input)
		if got != c.want {
			t.Errorf("WordFlip(%q) = %q, want %q", c.input, got, c.want)
		}
	}
}

// ─── 7. WEAREUNIQUE ──────────────────────────────────────────

func TestWeAreUnique(t *testing.T) {
	cases := []struct {
		s1, s2 string
		want   int
	}{
		{"foo", "boo", 2},  // f and b are unique
		{"", "", -1},
		{"abc", "def", 6},  // nothing shared
		{"abc", "abc", 0},  // everything shared
	}

	for _, c := range cases {
		got := WeAreUnique(c.s1, c.s2)
		if got != c.want {
			t.Errorf("WeAreUnique(%q, %q) = %d, want %d", c.s1, c.s2, got, c.want)
		}
	}
}

// ─── 8. ISCAPITALIZED ────────────────────────────────────────

func TestIsCapitalized(t *testing.T) {
	cases := []struct {
		input string
		want  bool
	}{
		{"Hello! How are you?", false}, // "are" and "you" start lowercase
		{"Hello How Are You", true},
		{"Whats 4this 100K?", true},    // non-alpha word start is OK
		{"Whatsthis4", true},
		{"!!!!Whatsthis4", true},
		{"", false},
		{"hello", false},
	}

	for _, c := range cases {
		got := IsCapitalized(c.input)
		if got != c.want {
			t.Errorf("IsCapitalized(%q) = %v, want %v", c.input, got, c.want)
		}
	}
}

// ─── 9. SAVEANDMISS ──────────────────────────────────────────

func TestSaveAndMiss(t *testing.T) {
	cases := []struct {
		arg  string
		num  int
		want string
	}{
		{"123456789", 3, "123789"},
		{"", 3, ""},
		{"hello", 0, "hello"},  // num=0 → return original
		{"hello", -1, "hello"}, // negative → return original
		{"ab", 5, "ab"},        // shorter than chunk → save all
	}

	for _, c := range cases {
		got := SaveAndMiss(c.arg, c.num)
		if got != c.want {
			t.Errorf("SaveAndMiss(%q, %d) = %q, want %q", c.arg, c.num, got, c.want)
		}
	}
}

// ─── 10. THIRDTIMEISACHARM ───────────────────────────────────

func TestThirdTimeIsACharm(t *testing.T) {
	cases := []struct {
		input string
		want  string
	}{
		{"123456789", "369\n"},
		{"", "\n"},
		{"12", "\n"},   // no third character
		{"abc", "c\n"},
		{"a b c d e f", "b e\n"}, // spaces are counted as positions
	}

	for _, c := range cases {
		got := ThirdTimeIsACharm(c.input)
		if got != c.want {
			t.Errorf("ThirdTimeIsACharm(%q) = %q, want %q", c.input, got, c.want)
		}
	}
}
