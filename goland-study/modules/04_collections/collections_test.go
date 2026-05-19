package collections

import "testing"

func TestSum(t *testing.T) {
	if got := Sum([]int{1, 2, 3, 4, 5}); got != 15 {
		t.Errorf("Sum([1..5]) = %d, want 15", got)
	}
	if got := Sum([]int{}); got != 0 {
		t.Errorf("Sum([]) = %d, want 0", got)
	}
	if got := Sum([]int{-1, 1}); got != 0 {
		t.Errorf("Sum([-1,1]) = %d, want 0", got)
	}
}

func TestCountWords(t *testing.T) {
	result := CountWords("go is go and go is great")
	if result["go"] != 3 {
		t.Errorf("Expected 'go' count = 3, got %d", result["go"])
	}
	if result["is"] != 2 {
		t.Errorf("Expected 'is' count = 2, got %d", result["is"])
	}
	if result["great"] != 1 {
		t.Errorf("Expected 'great' count = 1, got %d", result["great"])
	}
}

func TestFilter(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }
	result := Filter([]int{1, 2, 3, 4, 5, 6}, isEven)
	expected := []int{2, 4, 6}
	if len(result) != len(expected) {
		t.Fatalf("Filter length = %d, want %d", len(result), len(expected))
	}
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Filter[%d] = %d, want %d", i, result[i], v)
		}
	}
}
