## 6. Functions

### Quest 6.1: Function Factory
**Difficulty:** 🟡 Intermediate
**Skills:** Multiple returns, variadic, closures, recursion

```go
// solutions/06-functions/advanced.go
package main

import "fmt"

// Return multiple named values
func divide(dividend, divisor float64) (result float64, err error) {
    // TODO: Implement with proper error handling
    return
}

// Variadic function
func average(numbers ...float64) float64 {
    // TODO: Calculate average
    return 0
}

// Return a function (closure)
func makeMultiplier(factor int) func(int) int {
    // TODO: Return closure that multiplies by factor
    return nil
}

// Recursive function
func factorial(n int) int {
    // TODO: Implement recursively
    return 0
}

func main() {
    // Test all functions
}
```

**Tasks:**
- [ ] 1. Write functions that return multiple named values
- [ ] 2. Create a variadic function for calculating average
- [ ] 3. Implement closure that generates multiplier functions
- [ ] 4. Write recursive function for factorial
- [ ] 5. Create a function timer decorator that measures execution time
- [ ] 6. Implement memoization wrapper for expensive functions

### Quest 6.2: Error Handling Mastery
**Difficulty:** 🟡 Intermediate

```go
// solutions/06-functions/errors.go
package main

import "fmt"

// Custom error type
type DivisionError struct {
    Dividend float64
    Divisor  float64
    Message  string
}

func (e *DivisionError) Error() string {
    return fmt.Sprintf("cannot divide %f by %f: %s", 
        e.Dividend, e.Divisor, e.Message)
}

// Custom errors for different scenarios
type NegativeNumberError struct {
    Number float64
}

func (e *NegativeNumberError) Error() string {
    return fmt.Sprintf("negative number not allowed: %f", e.Number)
}

func safeSqrt(n float64) (float64, error) {
    if n < 0 {
        return 0, &NegativeNumberError{Number: n}
    }
    // Calculate sqrt
    return 0, nil
}

func main() {
    // Test error handling
    _, err := safeSqrt(-4)
    if err != nil {
        fmt.Println("Error:", err)
    }
}
```

**Tasks:**
- [ ] 1. Create custom error types for different scenarios (division, negative numbers, overflow)
- [ ] 2. Implement error wrapping with `fmt.Errorf()` and `%w` verb
- [ ] 3. Build retry logic with exponential backoff
- [ ] 4. Create an error chain inspector that unwraps and prints all errors
- [ ] 5. Implement a function that chains multiple error-prone operations

---
