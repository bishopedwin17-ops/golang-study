# Lecture 06: Pointers & Structs

---

## Pointers

A **pointer** holds the memory address of a value. Go uses pointers to allow functions to modify variables in the caller's scope.

```go
x := 42
p := &x         // p is a *int pointing to x
fmt.Println(*p) // 42 — dereference to get the value
*p = 100        // change the value at the address
fmt.Println(x)  // 100 — x changed!
```

### Why do pointers matter?
```go
// Without pointer — does NOT change the original
func doubleWrong(n int) {
    n *= 2 // operates on a copy
}

// With pointer — DOES change the original
func doubleRight(n *int) {
    *n *= 2
}

x := 5
doubleRight(&x)
fmt.Println(x) // 10
```

> **Rule of thumb**: Use pointers when you want a function to mutate the caller's variable, or when a value is very large and copying it would be expensive.

---

## Structs

Structs are custom types that group related fields together. This is Go's equivalent of a class (without the class keyword).

```go
type Person struct {
    Name string
    Age  int
    Email string
}

// Creating instances
alice := Person{Name: "Alice", Age: 30, Email: "alice@example.com"}
bob := Person{"Bob", 25, "bob@example.com"} // positional (fragile, avoid)

// Accessing fields
fmt.Println(alice.Name) // Alice
alice.Age = 31
```

### Pointers to Structs
```go
p := &Person{Name: "Charlie", Age: 22}
// Go auto-dereferences, so this works:
fmt.Println(p.Name) // Charlie (same as (*p).Name)
p.Age++             // Charlie is now 23
```

### Nested Structs
```go
type Address struct {
    Street string
    City   string
}

type Employee struct {
    Person  // embedded struct — "promotes" fields
    Address
    Company string
}

emp := Employee{
    Person:  Person{Name: "Dave", Age: 28},
    Address: Address{Street: "123 Main St", City: "Go City"},
    Company: "Gopherworks",
}

fmt.Println(emp.Name)   // promoted from Person
fmt.Println(emp.City)   // promoted from Address
```

### Anonymous (Struct Literal)
Useful for one-off data shapes:
```go
point := struct{ X, Y int }{X: 10, Y: 20}
```

---

## new() vs make()

| `new(T)`  | Allocates memory for T, returns `*T` with zero value |
|-----------|------------------------------------------------------|
| `make(T)` | Only for slices, maps, channels — returns initialized T (not a pointer) |
