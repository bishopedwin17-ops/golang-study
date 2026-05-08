# Lecture 01: Hello, Go!

Welcome to your first Go lesson! Go (or Golang) was created at Google to solve problems of scale and complexity. It's known for its simplicity, efficiency, and incredible support for concurrency.

## The Basic Structure of a Go Program

Every Go file starts with a **package declaration**. 
- Programs that are meant to be executed (like a CLI or a server) must be in `package main`.
- The entry point of a Go program is the `main` function.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### Key Components:

1.  **`package main`**: Tells the Go compiler that this file should compile as an executable program rather than a shared library.
2.  **`import "fmt"`**: Short for "format". This package contains functions for formatting text, including printing to the console.
3.  **`func main()`**: This is the first function that runs when you execute your program. It takes no arguments and returns nothing.
4.  **`fmt.Println(...)`**: A function from the `fmt` package that prints a line of text to the console.

## Why Go?
- **Fast Compilation**: Go compiles to machine code extremely quickly.
- **Static Typing**: Errors are caught at compile time, not runtime.
- **Garbage Collected**: You don't have to manually manage memory (like in C++).
- **No Classes**: Go uses structs and interfaces for a simpler approach to Object-Oriented Programming.

Ready to write your first Go code? Move on to `instructions.md`!
