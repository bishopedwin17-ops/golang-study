package exercises
import "testing"
func TestIsPrime(t *testing.T) {
	if !IsPrime(5) { t.Errorf("5 is prime") }
	if IsPrime(4) { t.Errorf("4 is not prime") }
	if IsPrime(1) { t.Errorf("1 is not prime") }
}
