package exercises
import "testing"
func TestShapes(t *testing.T) {
	shapes := []Shape{Circle{Radius: 2}, Rectangle{Width: 2, Height: 3}}
	res := TotalArea(shapes)
	if res < 18 || res > 19 { t.Errorf("Expected ~18.56, got %f", res) }
}
