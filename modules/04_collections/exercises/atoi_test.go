package exercises
import "testing"
func TestAtoi(t *testing.T) {
	tests := map[string]int{"123":123, "-123":-123, "0":0, "12a3":0, "+123":123, "":0}
	for k, v := range tests {
		if Atoi(k) != v { t.Errorf("Atoi(%q) expected %d", k, v) }
	}
}
