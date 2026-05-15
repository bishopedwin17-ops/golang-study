# ⚡ Go Quick Reference — Piscine Cheat Sheet

One page of everything you need. No standard library. Pure Go patterns.

---

## 📦 Package & Imports

```go
package main          // entry point package

import (
    "fmt"             // fmt.Println, fmt.Printf, fmt.Sprintf
    "os"              // os.Args, os.Stdout
)
```

---

## 📝 Variable Declaration

```go
// Declare + assign (short form — only inside functions)
x := 42
name := "hello"
flag := true

// Declare with type (can be at package level)
var count int         // zero value = 0
var text string       // zero value = ""
var ok bool           // zero value = false

// Multiple assignment
a, b := 1, 2
a, b = b, a           // swap!
```

---

## 🔢 Data Types

| Type | Description | Example |
|------|-------------|---------|
| `int` | whole number (positive/negative) | `42`, `-7` |
| `uint` | unsigned (no negative) | `0`, `255` |
| `byte` | alias for `uint8` (0–255) | `'A'` = 65 |
| `rune` | alias for `int32` — Unicode char | `'é'`, `'a'` |
| `string` | sequence of bytes (text) | `"hello"` |
| `bool` | true or false | `true`, `false` |
| `float64` | decimal number | `3.14` |

---

## 🔤 String Operations (no imports needed)

```go
s := "hello"

len(s)              // number of BYTES (not characters!)
len([]rune(s))      // number of CHARACTERS (safe for Unicode)

s[0]                // first byte (type: byte)
string(s[0])        // first byte as string

// Concatenation
s2 := s + " world"  // "hello world"
s2 += "!"           // "hello world!"

// Convert between string and rune slice
runes := []rune(s)  // string → []rune
back := string(runes) // []rune → string

// Convert a rune to a character string
ch := 'A'
str := string(ch)   // "A"
```

---

## 🔁 Loops

```go
// Basic for loop
for i := 0; i < 10; i++ {
    // i goes 0,1,2,...,9
}

// While-style loop
for condition {
    // loops while condition is true
}

// Infinite loop (use "break" to exit)
for {
    break
}

// Range over a STRING → gives runes
for i, ch := range "hello" {
    // i = index (byte position), ch = rune character
}
for _, ch := range s { // ignore index with _
}

// Range over a SLICE
for i, v := range slice {
    // i = index, v = value
}

// Range over a MAP
for key, val := range myMap {
}
```

---

## 🧩 Slices

```go
// Create
var s []int                  // nil slice (empty)
s := []int{1, 2, 3}          // slice literal
s := make([]int, 5)          // slice of 5 zeros
s := make([]int, 0, 10)      // empty slice, capacity 10

// Add element to end
s = append(s, 42)
s = append(s, 1, 2, 3)       // multiple elements
s = append(s1, s2...)         // append all of s2

// Length and capacity
len(s)      // number of elements
cap(s)      // underlying array capacity

// Slicing (sub-slice)
s[2:]       // from index 2 to end
s[:4]       // from start to index 3 (4 exclusive)
s[1:3]      // index 1 and 2

// Copy
dst := make([]int, len(src))
copy(dst, src)
```

---

## 🗺️ Maps

```go
// Create
m := make(map[string]int)
m := map[string]int{"a": 1, "b": 2}  // with initial values

// Set
m["key"] = 42

// Get (returns zero value if missing — no crash!)
val := m["key"]       // val = 0 if "key" not present
val, ok := m["key"]   // ok = false if missing

// Check existence
if m["key"] != 0 { }           // works for int maps
if m[ch] { }                   // works for map[rune]bool
if !seen[ch] { seen[ch]=true } // "have I seen this?" pattern

// Delete
delete(m, "key")

// Loop
for k, v := range m { }
```

---

## 🔀 Conditionals

```go
// if / else if / else
if x > 0 {
} else if x < 0 {
} else {
}

// switch (cleaner than many if-else)
switch ch {
case 'a', 'e', 'i', 'o', 'u':
    // vowel
case ' ':
    // space
default:
    // everything else
}

// switch with no condition (like if-else chain)
switch {
case x > 100:
case x > 50:
default:
}
```

---

## ⚙️ Functions

```go
// Basic
func add(a, b int) int {
    return a + b
}

// Multiple return values
func divide(a, b int) (int, bool) {
    if b == 0 {
        return 0, false
    }
    return a / b, true
}

// Variadic (accepts any number of arguments)
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
sum(1, 2, 3)       // 6
sum(mySlice...)    // unpack a slice as arguments
```

---

## 🖥️ Command-Line Arguments

```go
// os.Args[0] = program name (always)
// os.Args[1] = first argument
// os.Args[2] = second argument

if len(os.Args) != 2 {      // checks for exactly 1 user argument
    fmt.Println()            // print just newline
    return
}

arg := os.Args[1]            // the user's argument (type: string)
```

---

## 🔑 Key Operators

| Operator | Meaning | Example |
|----------|---------|---------|
| `+` | add / concatenate | `1+2`, `"a"+"b"` |
| `-` | subtract / negate | `5-3`, `-n` |
| `*` | multiply | `3*4` |
| `/` | integer divide | `7/2 = 3` |
| `%` | modulo (remainder) | `7%2 = 1` |
| `==` | equal | `x == 0` |
| `!=` | not equal | `x != 0` |
| `<` `>` `<=` `>=` | comparisons | |
| `&&` | logical AND | `a && b` |
| `\|\|` | logical OR | `a \|\| b` |
| `!` | logical NOT | `!ok` |
| `++` / `--` | increment / decrement | `i++` |
| `:=` | declare and assign | `x := 5` |
| `+=` | add and assign | `x += 1` |

---

## 🧮 Character Arithmetic (Piscine Patterns)

```go
// Is digit?
ch >= '0' && ch <= '9'

// Is lowercase letter?
ch >= 'a' && ch <= 'z'

// Is uppercase letter?
ch >= 'A' && ch <= 'Z'

// Is letter?
(ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')

// Digit char → int value
digit := int(ch - '0')   // '5' - '0' = 53-48 = 5

// Int value → digit char
ch := rune('0' + digit)  // 0+'0'='0', 5+'0'='5'

// Lowercase → uppercase (add 32 to ASCII? No — subtract 32)
upper := ch - 32         // 'a'(97) - 32 = 'A'(65)

// Uppercase → lowercase
lower := ch + 32         // 'A'(65) + 32 = 'a'(97)
```

---

## 📚 Stack Pattern (for brackets, RPN)

```go
var stack []int          // empty stack

// PUSH
stack = append(stack, val)

// PEEK (look at top without removing)
top := stack[len(stack)-1]

// POP (remove top)
top := stack[len(stack)-1]
stack = stack[:len(stack)-1]

// Is empty?
len(stack) == 0
```

---

## 📤 Output Patterns

```go
fmt.Println(x)           // prints x + newline
fmt.Print(x)             // prints x (NO newline)
fmt.Printf("%d\n", n)    // formatted: %d=int, %s=string, %c=char, %02x=hex
fmt.Sprintf("...")        // same but returns string instead of printing

// Format verbs:
// %d   integer
// %s   string
// %c   character (rune)
// %v   any value (default format)
// %#v  Go-syntax representation (useful for slices)
// %02x lowercase hex, zero-padded to 2 digits
// %02d decimal, zero-padded to 2 digits
```
