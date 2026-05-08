package exercises
import "testing"
func TestTimeoutFibonacci(t *testing.T) {
	res, err := TimeoutFibonacci(1)
	if err != nil || res != 1 { t.Errorf("Failed quick run") }
}
