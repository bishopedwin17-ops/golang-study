package exercises
import "testing"
func TestRecursiveFactorial(t *testing.T) {
	if RecursiveFactorial(5) != 120 { t.Errorf("Failed for 5") }
	if RecursiveFactorial(0) != 1 { t.Errorf("Failed for 0") }
	if RecursiveFactorial(-1) != 0 { t.Errorf("Failed for -1") }
}
