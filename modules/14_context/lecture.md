# Lecture 14: Context

`context.Context` is how Go handles **cancellation, timeouts, and deadlines** across goroutines and API boundaries.

---

## Why Context?

Imagine a web request that spawns a database query and an API call. If the user cancels the request, you want to cancel everything downstream too.

---

## Creating Contexts

```go
import "context"

// Root contexts (the "top" of a tree):
ctx := context.Background() // always non-nil; use at program start
ctx := context.TODO()       // placeholder when you're not sure yet

// Derived contexts:
ctx, cancel := context.WithCancel(parent)
ctx, cancel := context.WithTimeout(parent, 5*time.Second)
ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
```

> **Always call `cancel()`** when you're done (usually with `defer cancel()`). This prevents goroutine leaks.

---

## Using Context

```go
func fetchData(ctx context.Context, url string) ([]byte, error) {
    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil { return nil, err }

    resp, err := http.DefaultClient.Do(req)
    if err != nil { return nil, err } // will return error if ctx cancelled
    defer resp.Body.Close()
    return io.ReadAll(resp.Body)
}

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    data, err := fetchData(ctx, "https://api.example.com")
    if err != nil {
        log.Println("Request failed or timed out:", err)
    }
}
```

---

## Checking for Cancellation in Goroutines

```go
func worker(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            fmt.Println("Worker cancelled:", ctx.Err())
            return
        default:
            // do work
            time.Sleep(100 * time.Millisecond)
        }
    }
}
```

---

## Storing Values in Context

```go
type key string
const requestIDKey key = "requestID"

// Store:
ctx = context.WithValue(ctx, requestIDKey, "abc-123")

// Retrieve:
id := ctx.Value(requestIDKey).(string)
```

> Only store **request-scoped** values in context (request IDs, auth tokens). Don't put database connections, loggers, or large data there.

---

## The Rule

> *Pass context as the first parameter, call it `ctx`, and **always** check `ctx.Err()` or `<-ctx.Done()` in long-running operations.*
