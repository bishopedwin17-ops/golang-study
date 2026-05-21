package exercises
import "testing"
func TestMapFrequency(t *testing.T) {
	res := MapFrequency([]string{"a", "b", "a"})
	if res["a"] != 2 || res["b"] != 1 { t.Errorf("MapFrequency failed") }
}
