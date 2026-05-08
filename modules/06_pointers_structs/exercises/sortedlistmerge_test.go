package exercises
import "testing"
func TestSortedListMerge(t *testing.T) {
	n1 := &NodeI{Data: 1, Next: &NodeI{Data: 3, Next: nil}}
	n2 := &NodeI{Data: 2, Next: &NodeI{Data: 4, Next: nil}}
	res := SortedListMerge(n1, n2)
	expected := []int{1, 2, 3, 4}
	curr := res
	for _, v := range expected {
		if curr == nil || curr.Data != v { t.Errorf("Merge failed"); return }
		curr = curr.Next
	}
}
