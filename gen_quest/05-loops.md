## 5. Loops

### Quest 5.1: Loop Laboratory
**Difficulty:** 🟢 Beginner
**Skills:** For loops, range, break, continue

```go
// solutions/05-loops/patterns.go
package main

import "fmt"

func main() {
    // Task 1: Print Fibonacci sequence (first 20 numbers)
    // Task 2: Print multiplication table (12x12)
    // Task 3: FizzBuzz with custom rules
    // Task 4: Print diamond pattern
}
```

**Tasks:**
- [ ] 1. Print the Fibonacci sequence up to the 20th term
- [ ] 2. Create a multiplication table (12x12) using nested loops
- [ ] 3. Implement FizzBuzz from 1-100 with proper formatting
- [ ] 4. Print a diamond pattern of asterisks with adjustable size

### Quest 5.2: Game Loops
**Difficulty:** 🟡 Intermediate

```go
// solutions/05-loops/games.go
package main

import (
    "bufio"
    "fmt"
    "math/rand"
    "os"
    "time"
)

func numberGuessingGame() {
    rand.Seed(time.Now().UnixNano())
    target := rand.Intn(100) + 1
    scanner := bufio.NewScanner(os.Stdin)
    attempts := 0
    
    fmt.Println("Guess the number (1-100)!")
    
    for {
        fmt.Print("Enter guess: ")
        scanner.Scan()
        // TODO: Implement game logic
    }
}

func main() {
    numberGuessingGame()
}
```

**Tasks:**
- [ ] 1. Complete the number guessing game with 7 attempts
- [ ] 2. Build a simple menu system with submenus that loops until "quit"
- [ ] 3. Implement a countdown timer that prints every second
- [ ] 4. Create a prime number finder using `continue` and `break`

---
