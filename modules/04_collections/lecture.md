# Lecture 04: Collections — Arrays, Slices & Maps

Go has three fundamental collection types. Understanding the difference between arrays and slices is critical.

---

## Arrays

Arrays have a **fixed size** defined at compile time. You can't grow or shrink them.
```go
var primes [5]int = [5]int{2, 3, 5, 7, 11}
// shorthand:
primes := [5]int{2, 3, 5, 7, 11}
// let the compiler count:
primes := [...]int{2, 3, 5, 7, 11}

fmt.Println(primes[0]) // 2
fmt.Println(len(primes)) // 5
```

> Arrays are value types — assigning one array to another **copies** all elements.

---

## Slices ⭐ (The Go Workhorse)

Slices are dynamic, flexible **views** over an array. You'll use slices 95% of the time.

```go
// Create a slice
fruits := []string{"apple", "banana", "cherry"}

// Append to a slice
fruits = append(fruits, "mango")

// Slice a slice (fruits[low:high], high is exclusive)
fmt.Println(fruits[1:3]) // [banana cherry]
fmt.Println(fruits[:2])  // [apple banana]
fmt.Println(fruits[2:])  // [cherry mango]

fmt.Println(len(fruits)) // 4 — current elements
fmt.Println(cap(fruits)) // >= 4 — underlying array capacity
```

### make() for pre-sized slices
```go
// make([]Type, length, capacity)
scores := make([]int, 5)    // [0 0 0 0 0]
scores[0] = 100
```

### Iterating
```go
for i, v := range fruits {
    fmt.Printf("%d: %s\n", i, v)
}
```

---

## Maps

Maps are key-value stores (like dictionaries in Python or objects in JS).
```go
// Declare
ages := map[string]int{
    "Alice": 30,
    "Bob":   25,
}

// Add or update
ages["Charlie"] = 28
ages["Alice"] = 31

// Read
fmt.Println(ages["Bob"]) // 25

// Check if key exists (the "comma ok" idiom)
age, ok := ages["Dave"]
if !ok {
    fmt.Println("Dave not found")
} else {
    fmt.Println("Dave is", age)
}

// Delete
delete(ages, "Bob")

// Iterate
for name, age := range ages {
    fmt.Printf("%s: %d\n", name, age)
}
```

> Reading a missing key returns the **zero value** (0 for int, "" for string). Always use the `ok` idiom when presence matters.

---

## Key Takeaways
| Feature    | Array          | Slice           | Map               |
|------------|---------------|-----------------|-------------------|
| Size       | Fixed         | Dynamic         | Dynamic           |
| Index by   | int           | int             | any comparable    |
| Zero value | [N]T{0,0...}  | nil             | nil               |
| Make with  | `[N]T{}`      | `[]T{}` / make  | `map[K]V{}` / make|
