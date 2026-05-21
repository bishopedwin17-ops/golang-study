package generics

import "testing"

func TestSum(t *testing.T) {
	if got := Sum([]int{1, 2, 3, 4, 5}); got != 15 {
		t.Errorf("Sum(ints) = %d, want 15", got)
	}
	if got := Sum([]float64{1.5, 2.5}); got != 4.0 {
		t.Errorf("Sum(floats) = %f, want 4.0", got)
	}
}

func TestContains(t *testing.T) {
	if !Contains([]string{"go", "rust", "python"}, "go") {
		t.Error("Contains should find 'go'")
	}
	if Contains([]string{"go", "rust"}, "java") {
		t.Error("Contains should not find 'java'")
	}
	if !Contains([]int{1, 2, 3}, 2) {
		t.Error("Contains should find 2")
	}
}

func TestMap(t *testing.T) {
	doubled := Map([]int{1, 2, 3}, func(x int) int { return x * 2 })
	expected := []int{2, 4, 6}
	for i, v := range expected {
		if doubled[i] != v {
			t.Errorf("Map[%d] = %d, want %d", i, doubled[i], v)
		}
	}

	lengths := Map([]string{"go", "rust"}, func(s string) int { return len(s) })
	if lengths[0] != 2 || lengths[1] != 4 {
		t.Errorf("Map string lengths = %v, want [2 4]", lengths)
	}
}

func TestStack(t *testing.T) {
	s := &Stack[int]{}
	if s.Len() != 0 {
		t.Error("new stack should have length 0")
	}

	s.Push(10)
	s.Push(20)
	s.Push(30)

	if s.Len() != 3 {
		t.Errorf("stack length = %d, want 3", s.Len())
	}

	v, ok := s.Pop()
	if !ok || v != 30 {
		t.Errorf("Pop() = %d, want 30", v)
	}

	_, ok = (&Stack[string]{}).Pop()
	if ok {
		t.Error("Pop on empty stack should return false")
	}
}
