package exercises
import "testing"
func TestSwap(t *testing.T) {
	a, b := 1, 2
	Swap(&a, &b)
	if a != 2 || b != 1 { t.Errorf("Swap failed") }
}
