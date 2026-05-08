# Lecture 07: Methods & Interfaces

This is where Go's design philosophy shines. Instead of inheritance, Go uses **interfaces** — and they're implicit.

---

## Methods

A method is a function with a **receiver** (the type it belongs to).

```go
type Circle struct {
    Radius float64
}

// Value receiver — gets a copy, cannot mutate
func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

// Pointer receiver — can mutate the original
func (c *Circle) Grow(factor float64) {
    c.Radius *= factor
}
```

---

## Interfaces

An interface defines a **set of method signatures**. Any type that implements those methods **automatically** satisfies the interface — no `implements` keyword needed.

```go
type Shape interface {
    Area() float64
    Perimeter() float64
}
```

Now any type with `Area()` and `Perimeter()` methods satisfies `Shape`:

```go
type Circle struct { Radius float64 }
func (c Circle) Area() float64      { return math.Pi * c.Radius * c.Radius }
func (c Circle) Perimeter() float64 { return 2 * math.Pi * c.Radius }

type Rectangle struct { Width, Height float64 }
func (r Rectangle) Area() float64      { return r.Width * r.Height }
func (r Rectangle) Perimeter() float64 { return 2 * (r.Width + r.Height) }

// Now you can use them polymorphically:
func printShape(s Shape) {
    fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

printShape(Circle{5})
printShape(Rectangle{3, 4})
```

---

## The Empty Interface: `any` (or `interface{}`)

`any` accepts a value of **any type**. Use it when you truly don't know the type in advance.

```go
func describe(v any) {
    fmt.Printf("Value: %v, Type: %T\n", v, v)
}
describe(42)       // Value: 42, Type: int
describe("hello")  // Value: hello, Type: string
```

---

## Type Assertions & Type Switches

When you have an `any` value and want the underlying concrete type:

```go
// Type assertion (panics if wrong):
var v any = "hello"
s := v.(string) // s = "hello"

// Safe type assertion:
s, ok := v.(string)
if ok {
    fmt.Println("It's a string:", s)
}

// Type switch — the idiomatic way:
func whatAmI(v any) {
    switch x := v.(type) {
    case int:
        fmt.Println("int:", x)
    case string:
        fmt.Println("string:", x)
    case bool:
        fmt.Println("bool:", x)
    default:
        fmt.Printf("unknown type: %T\n", x)
    }
}
```

---

## Stringer Interface (Fmt)

Implement the `fmt.Stringer` interface to control how your type prints:
```go
type Person struct { Name string; Age int }

func (p Person) String() string {
    return fmt.Sprintf("%s (%d)", p.Name, p.Age)
}

fmt.Println(Person{"Alice", 30}) // Alice (30)
```

---

## Key Design Principle
> *"Accept interfaces, return concrete types."*

This is one of the most important Go idioms. Functions that accept an interface are more flexible and testable.
