package exercises
import "testing"
func TestPrintNbrStd(t *testing.T) {
	if PrintNbrStd(123) != "123" { t.Errorf("Failed") }
}
