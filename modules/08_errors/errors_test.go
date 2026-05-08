package errors_exercise

import (
	"errors"
	"testing"
)

func TestSafeDivide(t *testing.T) {
	result, err := SafeDivide(10, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != 5.0 {
		t.Errorf("SafeDivide(10,2) = %f, want 5.0", result)
	}

	_, err = SafeDivide(10, 0)
	if !errors.Is(err, ErrDivisionByZero) {
		t.Errorf("SafeDivide(10,0) should return ErrDivisionByZero, got %v", err)
	}
}

func TestParseAge(t *testing.T) {
	age, err := ParseAge("25")
	if err != nil {
		t.Fatalf("unexpected error for valid age: %v", err)
	}
	if age != 25 {
		t.Errorf("ParseAge(\"25\") = %d, want 25", age)
	}

	_, err = ParseAge("-1")
	if err == nil {
		t.Error("ParseAge(\"-1\") should return an error")
	}
	var ae *AgeError
	if !errors.As(err, &ae) {
		t.Errorf("ParseAge(\"-1\") should return *AgeError, got %T", err)
	}

	_, err = ParseAge("200")
	if err == nil {
		t.Error("ParseAge(\"200\") should return an error")
	}
}

func TestWrapError(t *testing.T) {
	inner := errors.New("file not found")
	wrapped := WrapError("loading data", inner)
	if wrapped == nil {
		t.Fatal("WrapError should not return nil")
	}
	if !errors.Is(wrapped, inner) {
		t.Error("wrapped error should unwrap to inner error")
	}
}
