# Lecture 03: Control Flow

Go has three control flow structures: `if/else`, `switch`, and `for`. That's it. No `while`, no `do-while` — just these three, and they're powerful enough to do everything.

---

## if / else

```go
age := 18

if age >= 18 {
    fmt.Println("Adult")
} else if age >= 13 {
    fmt.Println("Teenager")
} else {
    fmt.Println("Child")
}
```

### Initializer in if
Go allows an "init statement" before the condition. The variable is scoped only to the if/else block:
```go
if score := getScore(); score >= 90 {
    fmt.Println("A grade")
} else {
    fmt.Println("Try harder:", score)
}
// score is not accessible here
```

---

## switch

`switch` in Go doesn't need `break` — it stops automatically after each case.
```go
day := "Monday"

switch day {
case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
    fmt.Println("Weekday")
case "Saturday", "Sunday":
    fmt.Println("Weekend")
default:
    fmt.Println("Unknown")
}
```

### Expressionless switch (like a cleaner if/else chain)
```go
score := 85
switch {
case score >= 90:
    fmt.Println("A")
case score >= 80:
    fmt.Println("B")
default:
    fmt.Println("C or below")
}
```

---

## for — Go's only loop

### Basic for (like a while loop)
```go
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}
```

### Classic for (init; condition; post)
```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

### Infinite loop
```go
for {
    fmt.Println("forever")
    break // use break to exit
}
```

### Ranging over a slice
```go
fruits := []string{"apple", "banana", "cherry"}
for index, value := range fruits {
    fmt.Printf("Index %d: %s\n", index, value)
}
```

Use `_` to ignore the index if you don't need it:
```go
for _, fruit := range fruits {
    fmt.Println(fruit)
}
```

---

## Key Takeaways
- Go has NO `while` keyword — use `for` with just a condition.
- `switch` cases do NOT fall through by default (unlike C/Java). Use `fallthrough` keyword if you need it.
- `if` and `switch` can have init statements — great for keeping scope tight.

Ready? Open `instructions.md`!
