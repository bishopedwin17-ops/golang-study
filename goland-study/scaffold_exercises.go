package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var fileSystem = map[string]string{
	"modules/01_hello/exercises/hello_piscine.go": `package exercises

import "github.com/01-edu/z01"

// HelloPiscine prints "Hello Piscine!" using only z01.PrintRune
func HelloPiscine() {
	// Your code here
}
`,
	"modules/01_hello/exercises/hello_piscine_test.go": `package exercises
import "testing"
func TestHelloPiscine(t *testing.T) { t.Skip("Not implemented") }
`,
	"modules/01_hello/exercises/README.md": `# 01 Hello Exercises
Implement the function to print "Hello Piscine!" strictly using z01.PrintRune.
`,

	"modules/02_variables/exercises/var_declarations.go": `package exercises

// Fix the variable declarations to make this file compile.
// func Variables() {
//   var a int = "hello"
//   b := 
// }
`,
	"modules/02_variables/exercises/README.md": `# 02 Variables Exercises
Fix the declarations.
`,

	"modules/03_control_flow/exercises/printstr.go": `package exercises

import "github.com/01-edu/z01"

func PrintStr(s string) {
	// Your code here
}
`,
	"modules/03_control_flow/exercises/printnbr.go": `package exercises

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	// Your code here
}
`,
	"modules/03_control_flow/exercises/printcomb.go": `package exercises

import "github.com/01-edu/z01"

func PrintComb() {
	// Your code here
}
`,
	"modules/03_control_flow/exercises/printcomb2.go": `package exercises

import "github.com/01-edu/z01"

func PrintComb2() {
	// Your code here
}
`,
	"modules/03_control_flow/exercises/control_flow_test.go": `package exercises
import "testing"
func TestPrintStr(t *testing.T) { t.Skip("Not implemented") }
func TestPrintNbr(t *testing.T) { t.Skip("Not implemented") }
func TestPrintComb(t *testing.T) { t.Skip("Not implemented") }
func TestPrintComb2(t *testing.T) { t.Skip("Not implemented") }
`,
	"modules/03_control_flow/exercises/README.md": `# 03 Control Flow Exercises
Includes Piscine exercises: PrintStr, PrintNbr, PrintComb, PrintComb2.
`,

	"modules/04_collections/exercises/strlen.go": `package exercises

func StrLen(s string) int {
	return 0
}
`,
	"modules/04_collections/exercises/strrev.go": `package exercises

func StrRev(s string) string {
	return ""
}
`,
	"modules/04_collections/exercises/atoi.go": `package exercises

func Atoi(s string) int {
	return 0
}
`,
	"modules/04_collections/exercises/collections_test.go": `package exercises
import "testing"
func TestStrLen(t *testing.T) { t.Skip("Not implemented") }
func TestStrRev(t *testing.T) { t.Skip("Not implemented") }
func TestAtoi(t *testing.T) { t.Skip("Not implemented") }
`,
	"modules/04_collections/exercises/README.md": `# 04 Collections Exercises
Includes Piscine exercises: StrLen, StrRev, Atoi.
`,

	"modules/05_functions/exercises/recursivefactorial.go": `package exercises

func RecursiveFactorial(nb int) int {
	return 0
}
`,
	"modules/05_functions/exercises/fibonacci.go": `package exercises

func Fibonacci(index int) int {
	return 0
}
`,
	"modules/05_functions/exercises/isprime.go": `package exercises

func IsPrime(nb int) bool {
	return false
}
`,
	"modules/05_functions/exercises/functions_test.go": `package exercises
import "testing"
func TestRecursiveFactorial(t *testing.T) { t.Skip("Not implemented") }
func TestFibonacci(t *testing.T) { t.Skip("Not implemented") }
func TestIsPrime(t *testing.T) { t.Skip("Not implemented") }
`,
	"modules/05_functions/exercises/README.md": `# 05 Functions Exercises
Includes Piscine exercises: RecursiveFactorial, Fibonacci, IsPrime.
`,

	"modules/06_pointers_structs/exercises/swap.go": `package exercises

func Swap(a *int, b *int) {
	// Your code here
}
`,
	"modules/06_pointers_structs/exercises/divmod.go": `package exercises

func DivMod(a int, b int, div *int, mod *int) {
	// Your code here
}
`,
	"modules/06_pointers_structs/exercises/pointers_test.go": `package exercises
import "testing"
func TestSwap(t *testing.T) { t.Skip("Not implemented") }
func TestDivMod(t *testing.T) { t.Skip("Not implemented") }
`,
	"modules/06_pointers_structs/exercises/README.md": `# 06 Pointers & Structs Exercises
Includes Piscine exercises: Swap, DivMod.
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
	fmt.Println("Scaffolding complete.")
}
