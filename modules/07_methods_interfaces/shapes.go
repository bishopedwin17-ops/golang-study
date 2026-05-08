package methodsinterfaces

import "math"

// Shape is an interface for 2D geometric shapes.
// TODO: Define Area() float64 and Perimeter() float64 methods in this interface.
type Shape interface {
	// Your code here
}

// Circle represents a circle.
type Circle struct {
	Radius float64
}

// TODO: Implement Area() for Circle. Formula: π * r²
func (c Circle) Area() float64 {
	// Your code here
	return 0
}

// TODO: Implement Perimeter() for Circle. Formula: 2 * π * r
func (c Circle) Perimeter() float64 {
	// Your code here
	return 0
}

// Triangle represents a right triangle.
type Triangle struct {
	Base   float64
	Height float64
	Hyp    float64 // hypotenuse, used for perimeter
}

// TODO: Implement Area() for Triangle. Formula: 0.5 * base * height
func (t Triangle) Area() float64 {
	// Your code here
	return 0
}

// TODO: Implement Perimeter() for Triangle. Formula: Base + Height + Hyp
func (t Triangle) Perimeter() float64 {
	// Your code here
	return 0
}

// TotalArea returns the sum of Area() for all shapes.
// TODO: Implement using the Shape interface.
func TotalArea(shapes []Shape) float64 {
	// Your code here
	return 0
}

// keep — suppress import error
var _ = math.Pi
