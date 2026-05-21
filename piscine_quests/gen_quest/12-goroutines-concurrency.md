## 12. Goroutines & Concurrency

### Quest 12.1: Concurrency Basics
**Difficulty:** 🔴 Advanced
**Skills:** Goroutines, WaitGroups, channels

```go
// solutions/12-goroutines/basics.go
package main

import (
    "fmt"
    "sync"
    "time"
)

// Launch multiple goroutines
func launchGoroutines() {
    var wg sync.WaitGroup
    
    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            fmt.Printf("Goroutine %d starting\n", id)
            time.Sleep(time.Millisecond * time.Duration(id*100))
            fmt.Printf("Goroutine %d done\n", id)
        }(i)
    }
    
    wg.Wait()
    fmt.Println("All goroutines completed")
}

// Worker pool pattern
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
        time.Sleep(time.Millisecond * 100)
        results <- job * 2
    }
}

func workerPool() {
    const numJobs = 10
    const numWorkers = 3
    
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)
    var wg sync.WaitGroup
    
    // Start workers
    for w := 1; w <= numWorkers; w++ {
        wg.Add(1)
        go worker(w, jobs, results, &wg)
    }
    
    // Send jobs
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs)
    
    // Wait for workers and close results
    go func() {
        wg.Wait()
        close(results)
    }()
    
    // Collect results
    for result := range results {
        fmt.Printf("Result: %d\n", result)
    }
}

func main() {
    launchGoroutines()
    workerPool()
}
```

**Tasks:**
- [ ] 1. Launch 10 goroutines that print with random delays
- [ ] 2. Use `sync.WaitGroup` for synchronization
- [ ] 3. Implement worker pool pattern
- [ ] 4. Create fan-out, fan-in pattern
- [ ] 5. Build a pipeline with 3 stages (generate -> process -> output)
- [ ] 6. Implement concurrent web crawler with depth limit

### Quest 12.2: Channel Mastery
**Difficulty:** 🔴 Advanced

```go
// solutions/12-goroutines/channels.go
package main

import (
    "fmt"
    "time"
)

// Channel directions
func sender(ch chan<- string) {
    ch <- "Hello from sender"
}

func receiver(ch <-chan string) {
    msg := <-ch
    fmt.Println("Received:", msg)
}

// Select with timeout
func selectWithTimeout() {
    ch := make(chan string)
    
    go func() {
        time.Sleep(2 * time.Second)
        ch <- "result"
    }()
    
    select {
    case res := <-ch:
        fmt.Println(res)
    case <-time.After(1 * time.Second):
        fmt.Println("Timeout!")
    }
}

// Fan-in pattern
func fanIn(ch1, ch2 <-chan string) <-chan string {
    ch := make(chan string)
    
    go func() {
        for {
            select {
            case msg := <-ch1:
                ch <- "From ch1: " + msg
            case msg := <-ch2:
                ch <- "From ch2: " + msg
            }
        }
    }()
    
    return ch
}

// Or-done pattern
func orDone(done <-chan struct{}, c <-chan interface{}) <-chan interface{} {
    valStream := make(chan interface{})
    go func() {
        defer close(valStream)
        for {
            select {
            case <-done:
                return
            case v, ok := <-c:
                if !ok {
                    return
                }
                select {
                case valStream <- v:
                case <-done:
                }
            }
        }
    }()
    return valStream
}

func main() {
    // Unbuffered channel
    ch := make(chan string)
    go sender(ch)
    go receiver(ch)
    
    time.Sleep(time.Second)
    
    // Select with timeout
    selectWithTimeout()
}
```

**Tasks:**
- [ ] 1. Demonstrate channel directions (send-only, receive-only)
- [ ] 2. Implement select with timeout pattern
- [ ] 3. Build or-done channel pattern for cancellation
- [ ] 4. Create tee channel that duplicates messages to two receivers
- [ ] 5. Implement bridge channel that flattens channel of channels
- [ ] 6. Build a pub/sub system using channels

---
