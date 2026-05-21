## 14. Testing

### Quest 14.1: Test Suite Fundamentals
**Difficulty:** 🟡 Intermediate
**Skills:** Table-driven tests, benchmarks, examples

```go
// solutions/14-testing/math.go
package mathops

// Function to test
func Add(a, b int) int {
    return a + b
}

func Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}
```

```go
// solutions/14-testing/math_test.go
package mathops

import (
    "testing"
)

// Table-driven test
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive numbers", 1, 2, 3},
        {"negative numbers", -1, -2, -3},
        {"zero", 0, 0, 0},
        {"mixed", -1, 1, 0},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Add(%d, %d) = %d; want %d", 
                    tt.a, tt.b, result, tt.expected)
            }
        })
    }
}

// Test with subtests and helper
func TestDivide(t *testing.T) {
    t.Run("valid division", func(t *testing.T) {
        result, err := Divide(10, 2)
        if err != nil {
            t.Fatal("unexpected error:", err)
        }
        if result != 5 {
            t.Errorf("expected 5, got %f", result)
        }
    })
    
    t.Run("division by zero", func(t *testing.T) {
        _, err := Divide(10, 0)
        if err == nil {
            t.Error("expected error, got nil")
        }
    })
}

// Benchmark
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(1, 2)
    }
}

// Example test (appears in documentation)
func ExampleAdd() {
    result := Add(1, 2)
    fmt.Println(result)
    // Output: 3
}

// TestMain for setup/teardown
func TestMain(m *testing.M) {
    // Setup
    fmt.Println("Setting up tests...")
    
    // Run tests
    code := m.Run()
    
    // Teardown
    fmt.Println("Cleaning up...")
    
    os.Exit(code)
}
```

**Tasks:**
- [ ] 1. Write table-driven tests for calculator functions (Add, Subtract, Multiply, Divide)
- [ ] 2. Create benchmarks comparing string concatenation methods
- [ ] 3. Add example tests that appear in Go documentation
- [ ] 4. Achieve 80%+ test coverage and generate coverage report
- [ ] 5. Implement TestMain for setup and teardown
- [ ] 6. Organize tests with `t.Run()` for subtests

### Quest 14.2: Advanced Testing Patterns
**Difficulty:** 🔴 Advanced

**Tasks:**
- [ ] 1. Create mock HTTP server for testing HTTP clients
- [ ] 2. Implement database mock using interfaces
- [ ] 3. Use testify/assert for cleaner assertions
- [ ] 4. Implement golden file testing for complex outputs
- [ ] 5. Run tests in parallel with `t.Parallel()`
- [ ] 6. Implement fuzzing tests with `go test -fuzz`

---
