# Final Project: CLI Task Manager

Congratulations — you've made it to the final project! You will build a fully functional **command-line task manager** that puts everything you've learned to use.

---

## What You'll Build

A `tasks` CLI tool that can:
- `tasks add "Buy groceries"` — add a task
- `tasks list` — list all tasks
- `tasks done <id>` — mark a task complete
- `tasks delete <id>` — delete a task
- Tasks persist to a JSON file between runs

---

## Learning Goals Applied

| Concept | Where it appears |
|---------|-----------------|
| Structs | `Task` type |
| Interfaces | `Storage` interface |
| Error handling | File I/O, JSON, validation |
| JSON | Saving/loading tasks |
| os.Args / flag | CLI argument parsing |
| Packages | Organized project layout |
| Testing | Unit tests for all logic |

---

## Project Structure

```
modules/15_final_project/
├── main.go             ← entry point, CLI parsing
├── task/
│   ├── task.go         ← Task struct + TaskList type
│   └── task_test.go    ← tests for task logic
├── storage/
│   ├── storage.go      ← JSON file persistence
│   └── storage_test.go
└── instructions.md     ← step-by-step build guide
```

---

## Step-by-Step Instructions

### Step 1: Define the Task struct (in `task/task.go`)
```go
type Task struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    Done      bool      `json:"done"`
    CreatedAt time.Time `json:"created_at"`
}
```

### Step 2: Define a TaskList type with methods
- `Add(title string)` — adds a new task with auto-increment ID
- `Complete(id int) error` — marks task as done
- `Delete(id int) error` — removes task by ID
- `Pending() []Task` — returns only incomplete tasks

### Step 3: Implement JSON Storage (in `storage/storage.go`)
```go
type Storage interface {
    Load() ([]Task, error)
    Save(tasks []Task) error
}
```
Implement `JSONStorage` that reads/writes to a file.

### Step 4: Wire it all together in `main.go`
Parse `os.Args` and call the right methods.

### Step 5: Write tests
- Test `Add`, `Complete`, `Delete` with various edge cases.
- Test JSON round-trip (save then load, check equality).

---

## Bonus Challenges
- [ ] Add a `--file` flag to specify the data file path.
- [ ] Support `tasks list --all` to show completed tasks too.
- [ ] Add a `tasks clear` command.
- [ ] Add colored output using `\033[` ANSI escape codes.
- [ ] Sort tasks by creation date.

Good luck — you've got all the tools you need! 🚀
