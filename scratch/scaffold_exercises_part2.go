package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var fileSystem = map[string]string{
	"modules/07_methods_interfaces/exercises/shapes.go": `package exercises

// Shape interface and Circle/Rectangle structs.
`,
	"modules/07_methods_interfaces/exercises/README.md": `# 07 Methods & Interfaces Exercises
Implement Shape interface.
`,

	"modules/08_errors/exercises/safe_atoi.go": `package exercises

// SafeAtoi returns (int, error)
`,
	"modules/08_errors/exercises/README.md": `# 08 Errors Exercises
Implement a safe version of Atoi.
`,

	"modules/09_concurrency/exercises/async_print.go": `package exercises

// Print using goroutines.
`,
	"modules/09_concurrency/exercises/README.md": `# 09 Concurrency Exercises
Print strings concurrently.
`,

	"modules/10_testing/exercises/bad_atoi_test.go": `package exercises
import "testing"
// Write tests for a buggy Atoi.
`,
	"modules/10_testing/exercises/README.md": `# 10 Testing Exercises
Write tests for provided code.
`,

	"modules/11_standard_library/exercises/printnbr_stdlib.go": `package exercises

// Use fmt/strconv to print numbers.
`,
	"modules/11_standard_library/exercises/README.md": `# 11 Standard Library Exercises
Compare custom logic vs stdlib.
`,
}

func main() {
	for path, content := range fileSystem {
		dir := filepath.Dir(path)
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Printf("Error creating directory %s: %v\n", dir, err)
			continue
		}
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			fmt.Printf("Error writing file %s: %v\n", path, err)
			continue
		}
		fmt.Printf("Created: %s\n", path)
	}
	fmt.Println("Scaffolding Part 2 complete.")
}
