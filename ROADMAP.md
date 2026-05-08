# Go Programming Roadmap: From Zero to Pro

Welcome to your complete Go curriculum. Each module builds on the last — do them in order.

---

## ✅ Phase 1: The Building Blocks

| # | Module | Key Concepts | Status |
|---|--------|-------------|--------|
| 01 | [Hello, Go!](modules/01_hello/lecture.md) | main package, fmt, program structure | [ ] |
| 02 | [Variables & Types](modules/02_variables/lecture.md) | var, :=, const, zero values, basic types | [ ] |
| 03 | [Control Flow](modules/03_control_flow/lecture.md) | if/else, switch, for (Go's only loop) | [ ] |

## ✅ Phase 2: Data & Structure

| # | Module | Key Concepts | Status |
|---|--------|-------------|--------|
| 04 | [Collections](modules/04_collections/lecture.md) | Arrays, Slices, Maps, range | [ ] |
| 05 | [Functions](modules/05_functions/lecture.md) | Multiple returns, variadic, closures, defer | [ ] |
| 06 | [Pointers & Structs](modules/06_pointers_structs/lecture.md) | Memory addresses, custom types, embedding | [ ] |

## ✅ Phase 3: The Power of Go

| # | Module | Key Concepts | Status |
|---|--------|-------------|--------|
| 07 | [Methods & Interfaces](modules/07_methods_interfaces/lecture.md) | Receivers, implicit interfaces, type assertions | [ ] |
| 08 | [Error Handling](modules/08_errors/lecture.md) | error interface, wrapping, custom errors, panic | [ ] |
| 09 | [Concurrency](modules/09_concurrency/lecture.md) | Goroutines, channels, WaitGroup, Mutex, select | [ ] |

## ✅ Phase 4: Professional Go

| # | Module | Key Concepts | Status |
|---|--------|-------------|--------|
| 10 | [Testing](modules/10_testing/lecture.md) | Table-driven tests, subtests, benchmarks, -race | [ ] |
| 11 | [Standard Library](modules/11_standard_library/lecture.md) | fmt, strings, strconv, JSON, os, net/http, time | [ ] |
| 12 | [Generics](modules/12_generics/lecture.md) | Type parameters, constraints, generic data structures | [ ] |
| 13 | [Packages & Modules](modules/13_packages/lecture.md) | go.mod, exported names, internal, project layout | [ ] |
| 14 | [Context](modules/14_context/lecture.md) | Cancellation, timeouts, passing context through APIs | [ ] |

## ✅ Phase 5: Capstone

| # | Module | Key Concepts | Status |
|---|--------|-------------|--------|
| 15 | [Final Project: CLI Task Manager](modules/15_final_project/instructions.md) | Everything combined into a real tool | [ ] |

---

## How to Use This Course

1. **Read** the `lecture.md` in each module.
2. **Open** the `.go` file and fill in the `// TODO` sections.
3. **Test** your work: `cd modules/0X_name && go test -v`
4. **Move on** when all tests pass.

### Quick Commands
```bash
# Run tests for one module
go test -v ./modules/03_control_flow/...

# Run ALL tests at once
go test ./...

# Run with race detector (for module 09+)
go test -race ./...

# Run benchmarks (module 10+)
go test -bench=. ./modules/10_testing/...

# Build and run final project
go run ./modules/15_final_project/ add "Buy groceries"
go run ./modules/15_final_project/ list
```

---

## Mark Your Progress

When you complete a module, change `[ ]` to `[x]` in this file!
