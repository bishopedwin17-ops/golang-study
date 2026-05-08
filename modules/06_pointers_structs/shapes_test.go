package pointersstructs

import (
	"fmt"
	"testing"
)

func TestRectangleArea(t *testing.T) {
	r := Rectangle{Width: 5, Height: 3}
	if got := r.Area(); got != 15.0 {
		t.Errorf("Area() = %f, want 15.0", got)
	}
}

func TestRectangleScale(t *testing.T) {
	r := Rectangle{Width: 4, Height: 2}
	r.Scale(3)
	if r.Width != 12 {
		t.Errorf("After Scale(3), Width = %f, want 12", r.Width)
	}
	if r.Height != 6 {
		t.Errorf("After Scale(3), Height = %f, want 6", r.Height)
	}
}

func TestPersonGreet(t *testing.T) {
	p := Person{Name: "Alice", Age: 30}
	expected := fmt.Sprintf("Hi, I'm %s and I'm %d years old.", p.Name, p.Age)
	if got := p.Greet(); got != expected {
		t.Errorf("Greet() = %q, want %q", got, expected)
	}
}

func TestHaveBirthday(t *testing.T) {
	p := Person{Name: "Bob", Age: 25}
	p.HaveBirthday()
	if p.Age != 26 {
		t.Errorf("After HaveBirthday, Age = %d, want 26", p.Age)
	}
}
