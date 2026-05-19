## 7. Maps

### Quest 7.1: Map Mastery
**Difficulty:** 🟡 Intermediate
**Skills:** Maps, CRUD operations, word frequency

```go
// solutions/07-maps/frequency.go
package main

import (
    "fmt"
    "strings"
)

func wordFrequency(text string) map[string]int {
    // TODO: Count word occurrences
    return nil
}

func main() {
    text := "the quick brown fox jumps over the lazy dog the fox"
    freq := wordFrequency(text)
    
    for word, count := range freq {
        fmt.Printf("%s: %d\n", word, count)
    }
}
```

**Tasks:**
- [ ] 1. Implement word frequency counter from text string
- [ ] 2. Build a phone book with CRUD operations (Create, Read, Update, Delete)
- [ ] 3. Implement a cache with TTL (Time To Live) using maps
- [ ] 4. Create a URL shortener using nested maps (short code -> URL)

### Quest 7.2: Concurrent Safe Maps
**Difficulty:** 🔴 Advanced

```go
// solutions/07-maps/concurrent.go
package main

import (
    "fmt"
    "sync"
)

// Demonstrate race condition
func raceCondition() {
    m := make(map[int]int)
    // TODO: Show race condition with goroutines
}

// Fix with mutex
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

func main() {
    // Test implementations
}
```

**Tasks:**
- [ ] 1. Demonstrate a race condition with regular map and `go run -race`
- [ ] 2. Implement thread-safe map with `sync.RWMutex`
- [ ] 3. Use `sync.Map` for comparison
- [ ] 4. Benchmark all approaches with different workloads (read-heavy, write-heavy)
- [ ] 5. Create generic concurrent map wrapper using interfaces

---
