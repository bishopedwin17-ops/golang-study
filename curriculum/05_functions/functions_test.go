package functions

import (
	"testing"
)

func TestDivide(t *testing.T) {
	result, err := Divide(10, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != 5.0 {
		t.Errorf("Divide(10,2) = %f, want 5.0", result)
	}

	_, err = Divide(10, 0)
	if err == nil {
		t.Error("Divide(10,0) should return an error, got nil")
	}
}

func TestSumAll(t *testing.T) {
	if got := SumAll(1, 2, 3); got != 6 {
		t.Errorf("SumAll(1,2,3) = %d, want 6", got)
	}
	if got := SumAll(); got != 0 {
		t.Errorf("SumAll() = %d, want 0", got)
	}
}

func TestMakeAdder(t *testing.T) {
	add5 := MakeAdder(5)
	if add5 == nil {
		t.Fatal("MakeAdder returned nil")
	}
	if got := add5(3); got != 8 {
		t.Errorf("MakeAdder(5)(3) = %d, want 8", got)
	}
	if got := add5(0); got != 5 {
		t.Errorf("MakeAdder(5)(0) = %d, want 5", got)
	}
}

func TestMinMax(t *testing.T) {
	min, max := MinMax([]int{3, 1, 4, 1, 5, 9, 2, 6})
	if min != 1 {
		t.Errorf("MinMax min = %d, want 1", min)
	}
	if max != 9 {
		t.Errorf("MinMax max = %d, want 9", max)
	}
}
