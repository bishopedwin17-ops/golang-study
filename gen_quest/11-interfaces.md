## 11. Interfaces

### Quest 11.1: Interface Implementation
**Difficulty:** 🟡 Intermediate
**Skills:** Interfaces, polymorphism, type assertions

```go
// solutions/11-interfaces/shapes.go
package main

import (
    "fmt"
    "math"
)

// Shape interface
type Shape interface {
    Area() float64
    Perimeter() float64
}

// Circle implementation
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}

// Rectangle implementation
type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

// Triangle implementation
type Triangle struct {
    A, B, C float64 // sides
}

func (t Triangle) Area() float64 {
    // Heron's formula
    s := t.Perimeter() / 2
    return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}

func (t Triangle) Perimeter() float64 {
    return t.A + t.B + t.C
}

// Calculate total area of shapes
func totalArea(shapes []Shape) float64 {
    total := 0.0
    for _, shape := range shapes {
        total += shape.Area()
    }
    return total
}

// Type switch demonstration
func describeShape(s Shape) {
    fmt.Printf("Shape: %T\n", s)
    switch v := s.(type) {
    case Circle:
        fmt.Printf("Circle with radius %.2f\n", v.Radius)
    case Rectangle:
        fmt.Printf("Rectangle %.2f x %.2f\n", v.Width, v.Height)
    default:
        fmt.Printf("Unknown shape\n")
    }
}

func main() {
    shapes := []Shape{
        Circle{Radius: 5},
        Rectangle{Width: 4, Height: 6},
        Triangle{A: 3, B: 4, C: 5},
    }
    
    for _, shape := range shapes {
        fmt.Printf("Area: %.2f, Perimeter: %.2f\n", 
            shape.Area(), shape.Perimeter())
        describeShape(shape)
    }
    
    fmt.Printf("Total Area: %.2f\n", totalArea(shapes))
}
```

**Tasks:**
- [ ] 1. Create `Shape` interface with `Area()` and `Perimeter()` methods
- [ ] 2. Implement for: Circle, Rectangle, Triangle, Trapezoid
- [ ] 3. Add 3D shapes with `Volume()` method
- [ ] 4. Create function that calculates total area of shape slice
- [ ] 5. Implement empty interface demonstration
- [ ] 6. Practice type assertion and type switch patterns

### Quest 11.2: Advanced Interface Patterns
**Difficulty:** 🔴 Advanced

```go
// solutions/11-interfaces/sort.go
package main

import (
    "fmt"
    "sort"
)

// Custom type for sorting
type Person struct {
    Name string
    Age  int
}

// Implement sort.Interface
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByName []Person

func (n ByName) Len() int           { return len(n) }
func (n ByName) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n ByName) Less(i, j int) bool { return n[i].Name < n[j].Name }

func main() {
    people := []Person{
        {"Alice", 30},
        {"Bob", 25},
        {"Charlie", 35},
    }
    
    sort.Sort(ByAge(people))
    fmt.Println("Sorted by age:", people)
    
    sort.Sort(ByName(people))
    fmt.Println("Sorted by name:", people)
}
```

**Tasks:**
- [ ] 1. Implement `sort.Interface` for custom types
- [ ] 2. Create `io.Reader` and `io.Writer` wrappers with logging
- [ ] 3. Build a plugin system using interfaces
- [ ] 4. Implement strategy pattern for different sorting algorithms
- [ ] 5. Create middleware chain using `http.Handler` interface
- [ ] 6. Build dependency injection container

---
