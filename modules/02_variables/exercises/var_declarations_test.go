package exercises
import "testing"
func TestVariables(t *testing.T) {
	i, s, b, f := Variables()
	if i != 42 || s != "golang" || b != true || f != 3.14 {
		t.Errorf("Expected 42, 'golang', true, 3.14. Got %v, %v, %v, %v", i, s, b, f)
	}
}
