package exercises
import "testing"
func TestAsyncPrint(t *testing.T) {
	ch := AsyncPrint([]string{"A", "B", "C"})
	count := 0
	if ch != nil {
		for range ch { count++ }
		if count != 3 { t.Errorf("Did not receive all messages") }
	}
}
