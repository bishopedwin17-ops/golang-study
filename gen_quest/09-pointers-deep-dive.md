## 9. Pointers Deep Dive

### Quest 9.1: Pointer Power
**Difficulty:** 🔴 Advanced
**Skills:** Pointers, memory addresses, pass-by-reference

```go
// solutions/09-pointers/basics.go
package main

import "fmt"

// Swap two integers without temporary variable
func swap(a, b *int) {
    *a = *a + *b
    *b = *a - *b
    *a = *a - *b
}

// Change string via pointer
func appendGopher(s *string) {
    *s = *s + " Gopher!"
}

// Linked List Node
type Node struct {
    Value int
    Next  *Node
}

// Reverse linked list in place
func reverseList(head **Node) {
    var prev *Node
    current := *head
    
    for current != nil {
        next := current.Next
        current.Next = prev
        prev = current
        current = next
    }
    *head = prev
}

func main() {
    // Demonstrate swap
    x, y := 10, 20
    fmt.Printf("Before swap: x=%d, y=%d\n", x, y)
    swap(&x, &y)
    fmt.Printf("After swap: x=%d, y=%d\n", x, y)
    
    // Demonstrate string modification
    name := "Hello"
    appendGopher(&name)
    fmt.Println(name)  // "Hello Gopher!"
    
    // Explain pass-by-value
    z := 0
    setTo100(z)  // What happens here?
    fmt.Println(z)  // Still 0! Why?
}

func setTo100(y int) {
    y = 100  // This modifies copy, not original
}
```

**Tasks:**
- [ ] 1. Swap two variables without using a temporary variable (pointer arithmetic)
- [ ] 2. Create a function that modifies a string via pointer
- [ ] 3. Build a singly linked list with proper pointer management
- [ ] 4. Implement reverse linked list in place
- [ ] 5. Detect cycle in linked list (Floyd's algorithm)
- [ ] 6. Explain why `setTo100(z)` doesn't change `z`

### Quest 9.2: Memory Management & Escape Analysis
**Difficulty:** 🔴 Advanced

```go
// solutions/09-pointers/escape.go
package main

import "fmt"

// Stays on stack
func stackAllocation() int {
    x := 42
    return x
}

// Escapes to heap (returned pointer)
func heapAllocation() *int {
    y := 100
    return &y  // &y escapes to heap
}

// Build escape analysis with:
// go build -gcflags="-m" escape.go

func main() {
    a := stackAllocation()
    b := heapAllocation()
    fmt.Println(a, *b)
}
```

**Tasks:**
- [ ] 1. Demonstrate stack vs heap allocation
- [ ] 2. Use `go build -gcflags="-m"` to see escape analysis
- [ ] 3. Create a memory leak by losing pointer references
- [ ] 4. Profile memory with `runtime.ReadMemStats()`
- [ ] 5. Implement object pool pattern using `sync.Pool`
- [ ] 6. Profile with pprof to visualize heap usage

---
