## 1. Introduction & Setup

### Quest 1.1: First Contact
**Difficulty:** 🟢 Beginner
**Skills:** Installation, first program, workspace setup

```go
// solutions/01-intro/main.go
package main

import (
    "fmt"
    "runtime"
)

func main() {
    // TODO: Complete the following tasks
}
```

**Tasks:**
- [ ] 1. Install Go and verify with `go version`
- [ ] 2. Create module: `go mod init github.com/YOUR_USERNAME/golang-quests`
- [ ] 3. Print "Gopher initialized!" in green text using ANSI codes: `\033[32m`
- [ ] 4. Print your OS and architecture using `runtime.GOOS` and `runtime.GOARCH`
- [ ] 5. Build a binary with `go build` and compare file sizes (with/without debug symbols using `-ldflags="-s -w"`)
- [ ] 6. Measure execution time using the `time` command in terminal

**Challenge:** Write a program that detects and prints "Running on [OS Name]" with appropriate emoji (🐧 Linux, 🍎 macOS, 🪟 Windows)

### Quest 1.2: Project Structure
**Difficulty:** 🟢 Beginner

```
go-learning/
├── 01-intro/
│   ├── main.go
│   └── README.md
├── 02-variables/
├── go.mod
└── README.md
```

**Tasks:**
- [ ] 1. Create the folder structure above
- [ ] 2. Write a bash script that automates project setup
- [ ] 3. Add a `.gitignore` file for Go projects
- [ ] 4. Initialize git repository with meaningful commit messages

---
