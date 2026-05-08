package exercises
import "testing"
func TestFibonacci(t *testing.T) {
	if Fibonacci(4) != 3 { t.Errorf("Failed for 4") }
	if Fibonacci(-1) != -1 { t.Errorf("Failed for -1") }
}
