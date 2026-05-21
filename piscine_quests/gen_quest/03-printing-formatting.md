## 3. Printing & Formatting

### Quest 3.1: Format Master
**Difficulty:** 🟢 Beginner
**Skills:** fmt.Printf verbs, string formatting, aligned output

```go
// solutions/03-formatting/receipt.go
package main

import (
    "fmt"
    "time"
)

func main() {
    // Create a receipt that outputs:
    // ================================
    // ITEM          QTY    PRICE
    // ----          ---    -----
    // Laptop         1   $999.99
    // Mouse          2    $24.99
    // Keyboard       1    $79.99
    // ================================
    // TOTAL:              $1104.97
    // Date: 2024-01-15 14:30:00
    
    // TODO: Implement receipt printer
}
```

**Tasks:**
- [ ] 1. Use width specifiers for alignment (`%10s`, `%3d`, `%8.2f`)
- [ ] 2. Format prices to 2 decimal places
- [ ] 3. Add timestamp to receipt using `time.Now()`
- [ ] 4. Use every formatting verb at least once: `%v, %+v, %#v, %T, %t, %d, %b, %c, %x, %f, %s, %q`
- [ ] 5. Create both console and file output versions

### Quest 3.2: String Art & Tables
**Difficulty:** 🟡 Intermediate

**Tasks:**
- [ ] 1. Create a multiplication table (12x12) with perfect alignment
- [ ] 2. Build a progress bar formatter that shows percentage complete
- [ ] 3. Implement word wrap for terminal width (80 characters)
- [ ] 4. Create a function that takes any struct and prints it as a formatted table

**Challenge:** Build a function that takes a sentence, reverses word order, capitalizes alternate letters, and truncates/pads to exactly 50 characters.

---
