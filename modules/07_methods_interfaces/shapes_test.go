package methodsinterfaces

import (
	"math"
	"testing"
)

const eps = 1e-9

func TestCircleArea(t *testing.T) {
	c := Circle{Radius: 5}
	want := math.Pi * 25
	if got := c.Area(); math.Abs(got-want) > eps {
		t.Errorf("Circle.Area() = %f, want %f", got, want)
	}
}

func TestCirclePerimeter(t *testing.T) {
	c := Circle{Radius: 3}
	want := 2 * math.Pi * 3
	if got := c.Perimeter(); math.Abs(got-want) > eps {
		t.Errorf("Circle.Perimeter() = %f, want %f", got, want)
	}
}

func TestTriangleArea(t *testing.T) {
	tri := Triangle{Base: 6, Height: 4, Hyp: 0}
	if got := tri.Area(); math.Abs(got-12.0) > eps {
		t.Errorf("Triangle.Area() = %f, want 12.0", got)
	}
}

func TestTotalArea(t *testing.T) {
	shapes := []Shape{
		Circle{Radius: 1},
		Triangle{Base: 2, Height: 2},
	}
	want := math.Pi + 2.0
	if got := TotalArea(shapes); math.Abs(got-want) > eps {
		t.Errorf("TotalArea() = %f, want %f", got, want)
	}
}
