# Lecture 09: Concurrency — Goroutines & Channels

This is Go's killer feature. Go was designed from the ground up to make concurrency easy and safe.

---

## Goroutines

A goroutine is a lightweight thread managed by the Go runtime. Start one with the `go` keyword.

```go
func sayHello(name string) {
    fmt.Println("Hello,", name)
}

func main() {
    go sayHello("Alice") // runs concurrently
    go sayHello("Bob")   // runs concurrently
    time.Sleep(100 * time.Millisecond) // wait for goroutines (naive approach)
}
```

> Goroutines are cheap — you can run thousands of them. They're multiplexed onto OS threads by the Go scheduler.

---

## Channels

Channels allow goroutines to **communicate safely** by passing values.

```go
ch := make(chan int)     // unbuffered channel
ch := make(chan int, 5)  // buffered channel (holds up to 5 values)

// Send
ch <- 42

// Receive
value := <-ch
```

### Classic producer/consumer pattern
```go
func producer(ch chan<- int) { // chan<- means send-only
    for i := 0; i < 5; i++ {
        ch <- i
    }
    close(ch) // signal that no more values will be sent
}

func main() {
    ch := make(chan int)
    go producer(ch)

    for v := range ch { // range drains until channel is closed
        fmt.Println(v)
    }
}
```

---

## sync.WaitGroup — The Proper Way to Wait

```go
import "sync"

var wg sync.WaitGroup

for i := 0; i < 5; i++ {
    wg.Add(1) // increment counter before launching
    go func(id int) {
        defer wg.Done() // decrement when done
        fmt.Println("Worker", id)
    }(i)
}

wg.Wait() // blocks until counter reaches 0
```

---

## select — Non-blocking Channel Operations

`select` lets a goroutine wait on multiple channel operations simultaneously:

```go
select {
case msg := <-ch1:
    fmt.Println("from ch1:", msg)
case msg := <-ch2:
    fmt.Println("from ch2:", msg)
case <-time.After(1 * time.Second):
    fmt.Println("timeout!")
}
```

---

## sync.Mutex — Protecting Shared State

When goroutines share data (not via channels), use a `Mutex` to prevent race conditions:

```go
import "sync"

type SafeCounter struct {
    mu    sync.Mutex
    count int
}

func (c *SafeCounter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.count++
}

func (c *SafeCounter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.count
}
```

---

## The Go Concurrency Mantra
> *"Don't communicate by sharing memory; share memory by communicating."*

Prefer channels over shared state + mutexes when you can.

---

## Common Pitfalls
1. **Goroutine leak**: Always ensure goroutines can exit. Close channels or use `context.Context`.
2. **Closing a closed channel**: Panics. Only the sender should close.
3. **Race conditions**: Use `go test -race` to detect them.
