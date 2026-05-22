# Go Study Curriculum

Welcome to the Go Programming Roadmap repository. This project is designed as a complete, structured curriculum to take you from a beginner to a proficient Go developer through hands-on practice.

## Repository Structure

The curriculum is divided into a systematic, 15-module progression, starting with basic syntax and moving towards advanced topics like concurrency, standard libraries, and generics, culminating in a capstone project.

- `modules/`: Contains the 15 core learning modules. Each module contains reading material (`lecture.md`) and practical assignments.
- `piscine_exercises/`: Contains strict algorithm and logic exercises designed to enforce low-level programming skills without relying heavily on standard libraries.
- `ROADMAP.md`: A detailed breakdown of all the phases, modules, and your progress.

## Getting Started

1. Check the `ROADMAP.md` file to see the recommended order of modules.
2. Navigate to the module you want to work on (e.g., `cd modules/01_hello`).
3. Read the `lecture.md` file located inside the module.
4. Open the accompanying Go source files and implement the required logic.

## Testing Your Code

The repository uses Go's built-in testing framework to verify your solutions. Each module contains test files that you must pass before moving on to the next section.

To run tests for a specific module:
```bash
go test -v ./modules/01_hello/...
```

To run all tests across the entire repository:
```bash
go test ./...
```

For concurrency and advanced modules, use the race detector to catch hidden race conditions:
```bash
go test -race ./...
```

## Textbook Reference

A companion textbook has been added to the repository as a study asset:
- `resources/textbooks/thegobook.pdf`
- `resources/textbooks/thegobook/README.md`
- `resources/textbooks/thegobook/toc.md`
- `resources/textbooks/thegobook/mapping.md`

Use these files to map textbook chapters directly to the repo topics and exercises.

## Prerequisites

- Go 1.18 or higher installed on your system.
- A text editor or IDE of your choice.
