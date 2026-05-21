## 13. Mutex & Synchronization

### Quest 13.1: Race Condition Hunter
**Difficulty:** 🔴 Advanced
**Skills:** Mutex, RWMutex, race detection

```go
// solutions/13-sync/bank.go
package main

import (
    "fmt"
    "sync"
    "time"
)

// Race condition example - DON'T DO THIS!
type UnsafeCounter struct {
    count int
}

func (c *UnsafeCounter) Increment() {
    c.count++  // Race condition!
}

func (c *UnsafeCounter) Value() int {
    return c.count  // Race condition!
}

// Safe with Mutex
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

// Read-heavy optimization with RWMutex
type SafeMap struct {
    mu sync.RWMutex
    m  map[string]int
}

func (sm *SafeMap) Get(key string) (int, bool) {
    sm.mu.RLock()
    defer sm.mu.RUnlock()
    val, ok := sm.m[key]
    return val, ok
}

func (sm *SafeMap) Set(key string, value int) {
    sm.mu.Lock()
    defer sm.mu.Unlock()
    sm.m[key] = value
}

// Bank account with transfers
type BankAccount struct {
    mu      sync.Mutex
    balance int
}

func (a *BankAccount) Deposit(amount int) {
    a.mu.Lock()
    defer a.mu.Unlock()
    a.balance += amount
}

func (a *BankAccount) Withdraw(amount int) error {
    a.mu.Lock()
    defer a.mu.Unlock()
    if a.balance < amount {
        return fmt.Errorf("insufficient funds")
    }
    a.balance -= amount
    return nil
}

func (a *BankAccount) Balance() int {
    a.mu.Lock()
    defer a.mu.Unlock()
    return a.balance
}

// Transfer between accounts (careful with deadlocks!)
func transfer(from, to *BankAccount, amount int) error {
    // Lock ordering to prevent deadlock
    if from == to {
        return nil
    }
    
    from.mu.Lock()
    defer from.mu.Unlock()
    
    to.mu.Lock()
    defer to.mu.Unlock()
    
    if from.balance < amount {
        return fmt.Errorf("insufficient funds")
    }
    
    from.balance -= amount
    to.balance += amount
    return nil
}

func main() {
    // Demonstrate race condition
    fmt.Println("Unsafe counter (run with -race):")
    unsafe := &UnsafeCounter{}
    var wg sync.WaitGroup
    
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            unsafe.Increment()
        }()
    }
    wg.Wait()
    fmt.Printf("Unsafe count: %d (should be 1000)\n", unsafe.Value())
    
    // Safe counter
    safe := &SafeCounter{}
    
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            safe.Increment()
        }()
    }
    wg.Wait()
    fmt.Printf("Safe count: %d (should be 1000)\n", safe.Value())
    
    // Bank account test
    acc1 := &BankAccount{balance: 100}
    acc2 := &BankAccount{balance: 50}
    
    transfer(acc1, acc2, 30)
    fmt.Printf("Account1: %d, Account2: %d\n", acc1.Balance(), acc2.Balance())
    
    // Run this program with: go run -race bank.go
}
```

**Tasks:**
- [ ] 1. Create and detect race conditions with `go run -race`
- [ ] 2. Implement thread-safe counter with Mutex
- [ ] 3. Build thread-safe map with RWMutex for read-heavy workloads
- [ ] 4. Implement bank transfer function that prevents deadlocks
- [ ] 5. Compare performance: Mutex vs RWMutex vs atomic operations

### Quest 13.2: Advanced Synchronization Patterns
**Difficulty:** 🔴 Advanced

```go
// solutions/13-sync/patterns.go
package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

// Semaphore pattern using channels
type Semaphore chan struct{}

func NewSemaphore(n int) Semaphore {
    return make(Semaphore, n)
}

func (s Semaphore) Acquire() {
    s <- struct{}{}
}

func (s Semaphore) Release() {
    <-s
}

// Rate limiter
type RateLimiter struct {
    ticker *time.Ticker
}

func NewRateLimiter(interval time.Duration) *RateLimiter {
    return &RateLimiter{
        ticker: time.NewTicker(interval),
    }
}

func (rl *RateLimiter) Wait() {
    <-rl.ticker.C
}

// Connection pool
type ConnectionPool struct {
    mu      sync.Mutex
    conns   chan interface{}
    factory func() interface{}
}

func NewConnectionPool(size int, factory func() interface{}) *ConnectionPool {
    pool := &ConnectionPool{
        conns:   make(chan interface{}, size),
        factory: factory,
    }
    
    for i := 0; i < size; i++ {
        pool.conns <- factory()
    }
    
    return pool
}

func (p *ConnectionPool) Get() interface{} {
    return <-p.conns
}

func (p *ConnectionPool) Put(conn interface{}) {
    p.conns <- conn
}

// Circuit breaker
type CircuitBreaker struct {
    mu            sync.Mutex
    failures      int
    lastFailure   time.Time
    threshold     int
    resetTimeout  time.Duration
}

func NewCircuitBreaker(threshold int, timeout time.Duration) *CircuitBreaker {
    return &CircuitBreaker{
        threshold:    threshold,
        resetTimeout: timeout,
    }
}

func (cb *CircuitBreaker) Execute(fn func() error) error {
    cb.mu.Lock()
    
    if cb.failures >= cb.threshold {
        if time.Since(cb.lastFailure) < cb.resetTimeout {
            cb.mu.Unlock()
            return fmt.Errorf("circuit breaker open")
        }
        cb.failures = 0  // Reset after timeout
    }
    cb.mu.Unlock()
    
    err := fn()
    
    cb.mu.Lock()
    if err != nil {
        cb.failures++
        cb.lastFailure = time.Now()
    } else {
        cb.failures = 0
    }
    cb.mu.Unlock()
    
    return err
}

// Context-aware worker
func contextWorker(ctx context.Context, id int, wg *sync.WaitGroup) {
    defer wg.Done()
    
    for {
        select {
        case <-ctx.Done():
            fmt.Printf("Worker %d: shutting down\n", id)
            return
        default:
            fmt.Printf("Worker %d: working...\n", id)
            time.Sleep(time.Second)
        }
    }
}

func main() {
    // Semaphore example
    sem := NewSemaphore(3)
    for i := 0; i < 10; i++ {
        go func(id int) {
            sem.Acquire()
            defer sem.Release()
            fmt.Printf("Semaphore worker %d\n", id)
            time.Sleep(time.Second)
        }(i)
    }
    
    // Context example
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()
    
    var wg sync.WaitGroup
    for i := 0; i < 3; i++ {
        wg.Add(1)
        go contextWorker(ctx, i, &wg)
    }
    wg.Wait()
    
    time.Sleep(time.Second)
}
```

**Tasks:**
- [ ] 1. Implement semaphore using buffered channels
- [ ] 2. Build a rate limiter with time.Ticker
- [ ] 3. Create connection pool for managing resources
- [ ] 4. Implement circuit breaker pattern for fault tolerance
- [ ] 5. Build context-aware workers that respect cancellation
- [ ] 6. Compare performance of different sync primitives

---
