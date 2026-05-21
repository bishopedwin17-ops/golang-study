package exercises
import "testing"
func TestStrLen(t *testing.T) {
	if StrLen("Hello") != 5 { t.Errorf("StrLen('Hello') failed") }
	if StrLen("Hello World!") != 12 { t.Errorf("StrLen('Hello World!') failed") }
	if StrLen("♥") != 1 { t.Errorf("StrLen('♥') failed. Wait, string length vs rune length?") }
}
