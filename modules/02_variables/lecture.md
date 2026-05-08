# Lecture 02: Variables & Types

In Go, variables are explicitly typed, and the compiler uses these types to ensure your code is safe.

## Declaring Variables

There are three main ways to declare variables in Go:

### 1. The `var` Keyword
Use this when you want to declare a variable and initialize it later, or if you want to be explicit about the type.
```go
var name string = "Gopher"
var age int
age = 10
```

### 2. Short Variable Declaration (`:=`)
This is the most common way to declare and initialize variables inside functions. Go will **infer** the type for you.
```go
message := "Learning Go is fun!" // inferred as string
count := 42                     // inferred as int
```

### 3. Multiple Declaration
```go
var x, y, z int = 1, 2, 3
```

## Basic Types
- `string`: Text ("hello")
- `int`, `int8`, `int16`, `int32`, `int64`: Integers
- `float32`, `float64`: Floating point numbers
- `bool`: Boolean (true/false)

## Constants
Constants are declared like variables, but with the `const` keyword. They cannot be changed once declared.
```go
const Pi = 3.14
```

## Zero Values
If you declare a variable without an initial value, Go gives it a "Zero Value":
- `int`: `0`
- `string`: `""`
- `bool`: `false`

Ready to practice? Check `instructions.md`!
