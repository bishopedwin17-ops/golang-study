## 4. Arrays & Slices

### Quest 4.1: Slice Surgery
**Difficulty:** 🟡 Intermediate
**Skills:** Slices, append, copy, capacity management

```go
// solutions/04-slices/operations.go
package main

import "fmt"

// Remove every 3rd element from the slice
func removeEveryThird(slice []int) []int {
    // TODO: Implement efficiently
    return nil
}

// Insert sum of neighbors between each pair
func insertSumBetween(slice []int) []int {
    // Example: [1, 2, 3] -> [1, 3, 2, 5, 3]
    // TODO: Implement
    return nil
}

// Reverse slice in place (no new allocation)
func reverseInPlace(slice []int) {
    // TODO: Implement O(1) space
}

func main() {
    // Test your implementations here
}
```

**Tasks:**
- [ ] 1. Create a slice of 10 random integers
- [ ] 2. Implement `removeEveryThird()` function
- [ ] 3. Implement `insertSumBetween()` function
- [ ] 4. Implement `reverseInPlace()` function
- [ ] 5. Write benchmarks comparing approaches
- [ ] 6. Handle edge cases (empty slices, single elements)

### Quest 4.2: Memory Detective
**Difficulty:** 🔴 Advanced

```go
// solutions/04-slices/memory.go
package main

import "fmt"

func main() {
    // Task: Track when slice capacity changes
    slice := make([]int, 0, 1)
    for i := 0; i < 10; i++ {
        slice = append(slice, i)
        fmt.Printf("Len: %d, Cap: %d\n", len(slice), cap(slice))
    }
    // Notice the capacity growth pattern!
}
```

**Tasks:**
- [ ] 1. Track and visualize slice capacity growth pattern
- [ ] 2. Create a "slice leak" demonstration (holding reference to large underlying array)
- [ ] 3. Fix the leak by copying with `copy()` or `append([]int{}, slice...)`
- [ ] 4. Implement a circular buffer using slices
- [ ] 5. Benchmark `append` vs pre-allocation with `make()` for 1M elements

---
