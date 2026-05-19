package pointersstructs

// Rectangle represents a 2D rectangle.
type Rectangle struct {
	Width  float64
	Height float64
}

// Area returns the area of the rectangle.
// TODO: Implement (Width * Height).
func (r Rectangle) Area() float64 {
	// Your code here
	return 0
}

// Scale multiplies both Width and Height by factor.
// This must use a POINTER receiver so the change is visible to the caller.
// TODO: Implement.
func (r *Rectangle) Scale(factor float64) {
	// Your code here
}

// Person represents a human.
type Person struct {
	Name string
	Age  int
}

// Greet returns a greeting string like "Hi, I'm Alice and I'm 30 years old."
// TODO: Implement.
func (p Person) Greet() string {
	// Your code here
	return ""
}

// HaveBirthday increments the person's age by 1 using a pointer receiver.
// TODO: Implement.
func (p *Person) HaveBirthday() {
	// Your code here
}
