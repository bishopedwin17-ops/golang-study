package exercises
import "testing"
func TestGenericFilter(t *testing.T) {
	res := Filter([]int{1, 2, 3, 4}, func(x int) bool { return x%2 == 0 })
	if len(res) != 2 || res[0] != 2 { t.Errorf("Failed") }
}
