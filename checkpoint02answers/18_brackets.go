// =============================================================
// EXERCISE: brackets
// QUESTION: A PROGRAM that takes any number of arguments.
//           For each argument, print "OK\n" if the brackets are
//           correctly matched, or "Error\n" otherwise.
//           Bracket pairs: () [] {}
//           All other characters are ignored.
//           No arguments → print nothing.
// =============================================================
//
// HOW IT WORKS — STACK ALGORITHM:
//   We use a stack (last-in, first-out) to track opening brackets.
//   - When we see an opening bracket ( [ { → PUSH it onto the stack.
//   - When we see a closing bracket ) ] } → POP from the stack and
//     check it matches the corresponding opener.
//   - At the end, if the stack is empty → OK. Otherwise → Error.

package main

import (
	"fmt"
	"os"
)

// checkBrackets returns true if the brackets in s are correctly matched.
func checkBrackets(s string) bool {

	// "stack" is a slice of runes used as a stack (LIFO = Last In First Out).
	// DATA TYPE: []rune — a slice of rune values.
	// We start with an empty stack.
	var stack []rune

	for _, ch := range s {

		// Is this an opening bracket? Push it onto the stack.
		if ch == '(' || ch == '[' || ch == '{' {
			// "append" adds ch to the END of the slice.
			// In a stack, "end" = "top".
			stack = append(stack, ch)

		} else if ch == ')' || ch == ']' || ch == '}' {
			// Closing bracket found.
			// First, if the stack is empty, there is no matching opener → Error.
			if len(stack) == 0 {
				return false
			}

			// Pop the top of the stack.
			// The top is the LAST element: stack[len(stack)-1].
			// OPERATOR: - subtraction (to get the last index)
			top := stack[len(stack)-1]
			// Remove the top by re-slicing: stack[:len-1] gives everything BUT the last.
			// DATA TYPE: []rune — slicing creates a new view of the same underlying array.
			stack = stack[:len(stack)-1]

			// Check that the popped opener matches the current closer.
			// A closing ')' must pair with '(', etc.
			if (ch == ')' && top != '(') ||
				(ch == ']' && top != '[') ||
				(ch == '}' && top != '{') {
				return false // mismatched pair
			}
		}
		// All other characters are ignored (no else branch needed).
	}

	// After processing all characters, the stack must be empty.
	// If there are leftover openers, they were never closed → Error.
	return len(stack) == 0
}

func main() {

	// os.Args[1:] is a slice containing all arguments AFTER the program name.
	// This gives us just the user-provided arguments.
	args := os.Args[1:]

	if len(args) == 0 {
		return // print nothing
	}

	for _, arg := range args {
		if checkBrackets(arg) {
			fmt.Println("OK")
		} else {
			fmt.Println("Error")
		}
	}
}

// ─── HOW TO RUN ──────────────────────────────────────────────
// go run 18_brackets.go "(johndoe)"            → OK
// go run 18_brackets.go "([)]"                 → Error
// go run 18_brackets.go "" "{[(0+0)(1+1)]{()}}"  → OK \n OK
// go run 18_brackets.go                        → (nothing)
