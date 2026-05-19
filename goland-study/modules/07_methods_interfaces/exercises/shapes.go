package exercises

type Shape interface { Area() float32 }
type Circle struct { Radius float32 }
type Rectangle struct { Width, Height float32 }
func (c Circle) Area() float32 { return 0 }
func (r Rectangle) Area() float32 { return 0 }
func TotalArea(shapes []Shape) float32 { return 0 }
