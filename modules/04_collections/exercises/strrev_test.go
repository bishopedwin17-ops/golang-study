package exercises
import "testing"
func TestStrRev(t *testing.T) {
	if StrRev("Hello") != "olleH" { t.Errorf("StrRev failed") }
}
