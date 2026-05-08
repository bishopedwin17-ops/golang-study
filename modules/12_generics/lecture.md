# Lecture 12: Generics (Go 1.18+)

Go added generics in 1.18. They let you write functions and types that work with any type while remaining type-safe.

---

## The Problem Without Generics

```go
// Without generics, you'd need one function per type:
func SumInts(nums []int) int { ... }
func SumFloat64s(nums []float64) float64 { ... }
```

---

## Type Parameters

```go
// [T int | float64] is a "type constraint"
func Sum[T int | float64](nums []T) T {
    var total T
    for _, n := range nums {
        total += n
    }
    return total
}

Sum([]int{1, 2, 3})         // 6
Sum([]float64{1.1, 2.2})    // 3.3
```

---

## Constraints with `comparable`

`comparable` means the type supports `==` and `!=`:
```go
func Contains[T comparable](slice []T, item T) bool {
    for _, v := range slice {
        if v == item {
            return true
        }
    }
    return false
}

Contains([]string{"go", "rust"}, "go")  // true
Contains([]int{1, 2, 3}, 4)             // false
```

---

## The `constraints` Package Pattern

Define your own constraints using interfaces:
```go
type Number interface {
    int | int8 | int16 | int32 | int64 |
    float32 | float64
}

func Min[T Number](a, b T) T {
    if a < b { return a }
    return b
}
```

---

## Generic Types (Data Structures)

```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
    var zero T
    if len(s.items) == 0 {
        return zero, false
    }
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item, true
}

// Usage:
s := &Stack[int]{}
s.Push(1)
s.Push(2)
v, _ := s.Pop() // 2
```

---

## When to Use Generics
✅ Writing reusable data structures (stacks, queues, sets)  
✅ Writing utility functions that work on multiple types  
❌ When interfaces already solve the problem cleanly  
❌ Don't reach for generics by default — only when you see real duplication
