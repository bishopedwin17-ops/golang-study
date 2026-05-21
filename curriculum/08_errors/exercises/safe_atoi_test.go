package exercises
import "testing"
func TestSafeAtoi(t *testing.T) {
	res, err := SafeAtoi("123")
	if err != nil || res != 123 { t.Errorf("Failed for valid") }
	_, err2 := SafeAtoi("12a3")
	if err2 == nil { t.Errorf("Failed to return error for invalid") }
}
