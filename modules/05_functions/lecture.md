# Lecture 05: Functions

Functions are first-class citizens in Go. They can be passed as arguments, returned from other functions, and assigned to variables.

---

## Basics

```go
func add(a int, b int) int {
    return a + b
}

// If params share type, you can write:
func add(a, b int) int {
    return a + b
}
```

---

## Multiple Return Values ⭐

This is one of Go's most important features. Go functions commonly return `(result, error)`.

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}

// Calling it:
result, err := divide(10, 2)
if err != nil {
    log.Fatal(err)
}
fmt.Println(result) // 5
```

---

## Named Return Values

You can name return values, which makes the `return` statement implicit (a "naked return"). Use sparingly — only when it improves clarity.

```go
func minMax(nums []int) (min, max int) {
    min, max = nums[0], nums[0]
    for _, n := range nums[1:] {
        if n < min { min = n }
        if n > max { max = n }
    }
    return // naked return — returns min and max
}
```

---

## Variadic Functions

A function that accepts any number of arguments of the same type.
```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

sum(1, 2)
sum(1, 2, 3, 4, 5)

// Expand a slice into variadic args with ...
nums := []int{1, 2, 3}
sum(nums...)
```

---

## First-Class Functions & Closures

Functions can be stored in variables and passed around:
```go
// Function assigned to variable
square := func(x int) int { return x * x }
fmt.Println(square(4)) // 16

// Passing functions as arguments (higher-order functions)
func apply(nums []int, f func(int) int) []int {
    result := make([]int, len(nums))
    for i, v := range nums {
        result[i] = f(v)
    }
    return result
}

doubled := apply([]int{1, 2, 3}, func(x int) int { return x * 2 })
```

### Closures
A closure "captures" variables from its surrounding scope:
```go
func makeCounter() func() int {
    count := 0
    return func() int {
        count++ // captures 'count'
        return count
    }
}

counter := makeCounter()
fmt.Println(counter()) // 1
fmt.Println(counter()) // 2
fmt.Println(counter()) // 3
```

---

## defer ⭐

`defer` schedules a function to run **after** the current function returns. Perfect for cleanup (closing files, releasing locks, etc.).

```go
func readFile(path string) {
    f, err := os.Open(path)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close() // runs when readFile() exits — guaranteed!

    // ... work with file
}
```

Multiple defers run in **LIFO** (last in, first out) order:
```go
defer fmt.Println("first deferred - runs last")
defer fmt.Println("second deferred - runs first")
```
