package piscine

import (
	"reflect"
	"testing"
)

// ─── 1. SPLIT TEST ───────────────────────────────────────────
func TestSplit(t *testing.T) {
	tests := []struct {
		s, sep string
		want   []string
	}{
		{"HelloHAhowHAareHAyou?", "HA", []string{"Hello", "how", "are", "you?"}},
		{"a|b|c", "|", []string{"a", "b", "c"}},
		{"no_separator_here", "SEP", []string{"no_separator_here"}},
	}
	for _, tc := range tests {
		got := Split(tc.s, tc.sep)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Split(%q, %q) = %q; want %q", tc.s, tc.sep, got, tc.want)
		}
	}
}

// ─── 2. ACTIVE BITS TEST ─────────────────────────────────────
func TestActiveBits(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{7, 3},  // 111
		{15, 4}, // 1111
		{2, 1},  // 10
		{0, 0},
	}
	for _, tc := range tests {
		got := ActiveBits(tc.n)
		if got != tc.want {
			t.Errorf("ActiveBits(%d) = %d; want %d", tc.n, got, tc.want)
		}
	}
}

// ─── 3. SORT WORD ARRAY TEST ─────────────────────────────────
func TestSortWordArr(t *testing.T) {
	arr := []string{"a", "A", "1", "b", "B", "2", "c", "C", "3"}
	want := []string{"1", "2", "3", "A", "B", "C", "a", "b", "c"}
	
	SortWordArr(arr)
	
	if !reflect.DeepEqual(arr, want) {
		t.Errorf("SortWordArr failed.\nGot:  %q\nWant: %q", arr, want)
	}
}

// ─── 4. LIST SIZE TEST ───────────────────────────────────────
func TestListSize(t *testing.T) {
	// Manually build a list to test ListSize without relying on ListPushBack
	n3 := &NodeL{Data: "third", Next: nil}
	n2 := &NodeL{Data: "second", Next: n3}
	n1 := &NodeL{Data: "first", Next: n2}
	l := &List{Head: n1, Tail: n3}

	got := ListSize(l)
	if got != 3 {
		t.Errorf("ListSize() = %d; want 3", got)
	}

	emptyList := &List{}
	if ListSize(emptyList) != 0 {
		t.Errorf("ListSize() on empty list = %d; want 0", ListSize(emptyList))
	}
}

// ─── 5. LIST LAST TEST ───────────────────────────────────────
func TestListLast(t *testing.T) {
	n3 := &NodeL{Data: "third", Next: nil}
	n2 := &NodeL{Data: "second", Next: n3}
	n1 := &NodeL{Data: "first", Next: n2}
	l := &List{Head: n1, Tail: n3}

	got := ListLast(l)
	if got != "third" {
		t.Errorf("ListLast() = %v; want 'third'", got)
	}

	emptyList := &List{}
	if ListLast(emptyList) != nil {
		t.Errorf("ListLast() on empty list should return nil")
	}
}
