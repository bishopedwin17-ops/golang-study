## 2. Variables & Types

### Quest 2.1: Type Explorer
**Difficulty:** 🟢 Beginner
**Skills:** Variable declaration, basic types, zero values, type conversion

```go
// solutions/02-variables/types.go
package main

import "fmt"

func main() {
    // Task 1: Declare variables using all 3 methods
    var name string = "Gopher"
    age := 10
    const Pi = 3.14159
    
    // Task 2: Print zero values for all basic types
    var (
        intVar int
        floatVar float64
        stringVar string
        boolVar bool
    )
    
    fmt.Printf("Zero values:\nint: %d\nfloat64: %f\nstring: %q\nbool: %t\n", 
        intVar, floatVar, stringVar, boolVar)
    
    // TODO: Complete remaining tasks
}
```

**Tasks:**
- [ ] 1. Declare variables using all three methods: `var`, `:=`, and `const`
- [ ] 2. Print zero values for: `int`, `float64`, `string`, `bool`, `slice`, `map`, `pointer`
- [ ] 3. Create a temperature converter (Celsius ↔ Fahrenheit) with proper `float64` types
- [ ] 4. Print the type of each variable using `%T` verb
- [ ] 5. Store formatted strings in variables using `fmt.Sprintf()`

**Bug Fix Challenge:**
```go
// This program has 5 type-related bugs. Find and fix them all.
// Add comments explaining each fix.
func buggyProgram() {
    var a int = 5.5           // Bug 1
    b := "10"
    c := a + b                // Bug 2
    var d float64 = "3.14"    // Bug 3
    e := int("42")            // Bug 4
    var f string = 42         // Bug 5
    fmt.Println(a, b, c, d, e, f)
}
```

### Quest 2.2: Type Conversion Master
**Difficulty:** 🟡 Intermediate

**Tasks:**
- [ ] 1. Parse numeric strings to integers safely using `strconv.Atoi()`
- [ ] 2. Implement safe division function that handles division by zero
- [ ] 3. Create a function that accepts `interface{}` and prints its dynamic type
- [ ] 4. Build a type registry that maps strings to types using reflection

---
