package exercises
import "testing"
func TestDivMod(t *testing.T) {
	var div, mod int
	DivMod(13, 2, &div, &mod)
	if div != 6 || mod != 1 { t.Errorf("DivMod failed") }
}
