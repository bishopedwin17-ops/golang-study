# Lecture 08: Error Handling

Go takes a different approach to errors than most languages. There's no `try/catch`. Instead, errors are just **values** — you return them and check them explicitly.

---

## The `error` Interface

`error` is a built-in interface with a single method:
```go
type error interface {
    Error() string
}
```

Any type that has an `Error() string` method satisfies this interface.

---

## Creating Errors

### Simple errors
```go
import "errors"
err := errors.New("something went wrong")
```

### Formatted errors
```go
import "fmt"
err := fmt.Errorf("user %d not found", userID)
```

---

## Returning and Checking Errors

The standard Go pattern is: **return `(result, error)`** and check `err != nil`.

```go
func openFile(name string) (string, error) {
    if name == "" {
        return "", errors.New("filename cannot be empty")
    }
    // ... read file
    return contents, nil
}

contents, err := openFile("data.txt")
if err != nil {
    log.Fatalf("failed to open file: %v", err)
}
fmt.Println(contents)
```

---

## Custom Error Types

For richer errors (with error codes, fields, etc.), create a struct that implements the `error` interface:

```go
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation failed on field %q: %s", e.Field, e.Message)
}

func validateAge(age int) error {
    if age < 0 {
        return &ValidationError{Field: "age", Message: "must be non-negative"}
    }
    return nil
}
```

---

## Wrapping Errors (`%w`)

Since Go 1.13, you can **wrap** errors to add context while preserving the original:

```go
err := fmt.Errorf("processing user: %w", originalErr)

// Unwrap to check the original:
if errors.Is(err, os.ErrNotExist) {
    fmt.Println("File doesn't exist")
}

// Unwrap to get typed error:
var ve *ValidationError
if errors.As(err, &ve) {
    fmt.Println("Validation field:", ve.Field)
}
```

---

## errors.Is vs errors.As

| Function | When to use |
|----------|------------|
| `errors.Is(err, target)` | Check if error **is** a specific sentinel value |
| `errors.As(err, &target)` | Check if error **is of a specific type** and extract it |

---

## panic & recover (use sparingly!)

`panic` stops normal execution — only use it for truly unrecoverable situations.

```go
func mustDivide(a, b int) int {
    if b == 0 {
        panic("division by zero — programmer error!")
    }
    return a / b
}
```

`recover` can catch a panic (e.g., in a web server to prevent a crash):
```go
func safeDiv(a, b int) (result int, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("recovered from panic: %v", r)
        }
    }()
    return a / b, nil
}
```

> **Rule**: Return errors for expected failures. Use `panic` only for programming errors (bugs), never for normal control flow.
