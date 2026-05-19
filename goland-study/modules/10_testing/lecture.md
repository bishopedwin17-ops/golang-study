# Lecture 10: Testing in Go

Go has a built-in testing framework — no external libraries needed.

---

## The Basics

- Test files end in `_test.go`.
- Test functions start with `Test` and take `*testing.T`.
- Run with `go test ./...` (all packages) or `go test -v` (verbose).

```go
package mypackage

import "testing"

func TestAdd(t *testing.T) {
    got := Add(2, 3)
    want := 5
    if got != want {
        t.Errorf("Add(2, 3) = %d, want %d", got, want)
    }
}
```

---

## Table-Driven Tests ⭐ (The Go Standard)

Instead of writing one test per case, use a table:

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name string
        a, b int
        want int
    }{
        {"positive", 2, 3, 5},
        {"zeros", 0, 0, 0},
        {"negative", -1, -2, -3},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := Add(tt.a, tt.b)
            if got != tt.want {
                t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
            }
        })
    }
}
```

---

## Subtests with t.Run

`t.Run` creates a named subtest. You can run individual subtests:
```
go test -run TestAdd/negative
```

---

## Benchmarks

Benchmarks start with `Benchmark` and take `*testing.B`:

```go
func BenchmarkFibonacci(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Fibonacci(20)
    }
}
```

Run with: `go test -bench=.`

---

## Test Helpers

Use `t.Helper()` in helper functions so error lines point to the caller:
```go
func assertEqual(t *testing.T, got, want int) {
    t.Helper() // marks this as a helper
    if got != want {
        t.Errorf("got %d, want %d", got, want)
    }
}
```

---

## Race Detector

```
go test -race ./...
```
This detects concurrent data races. Always run before shipping.

---

## Coverage

```
go test -cover ./...
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

Aim for 80%+ on business-critical code, not 100% on everything.

---

## Testing HTTP Handlers

Use `net/http/httptest`:
```go
func TestHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/hello", nil)
    w := httptest.NewRecorder()
    HelloHandler(w, req)

    resp := w.Result()
    if resp.StatusCode != http.StatusOK {
        t.Errorf("status = %d, want %d", resp.StatusCode, http.StatusOK)
    }
}
```
