// =============================================================
// EXERCISE: rpncalc
// QUESTION: A PROGRAM that evaluates a Reverse Polish Notation (RPN)
//           expression given as a single string argument.
//           Operators: + - * / %
//           Print result + \n, or "Error\n" if invalid or wrong args.
// =============================================================
//
// WHAT IS RPN?
//   Normal:  3 + 4        (operator between operands)
//   RPN:     3 4 +        (operator AFTER operands)
//
//   "1 2 * 3 * 4 +" means:
//     1 2 *  → 1×2 = 2
//     2 3 *  → 2×3 = 6
//     6 4 +  → 6+4 = 10
//
// ALGORITHM: Stack-based evaluation.
//   - Numbers → PUSH onto stack.
//   - Operators → POP two numbers, apply operator, PUSH result.
//   - At the end, one number left on stack = answer.

package main

import (
	"fmt"
	"os"
)

// isDigitOrSign returns true if byte b starts a number token.
func isDigitStart(b byte) bool {
	return b >= '0' && b <= '9'
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Error")
		return
	}

	s := os.Args[1]

	// "stack" holds integer operands as we process them.
	// DATA TYPE: []int — a slice of ints used as a stack.
	var stack []int

	// Walk through the string token by token.
	// Tokens are separated by spaces.
	i := 0
	n := len(s)

	for i < n {

		// Skip whitespace (spaces).
		for i < n && s[i] == ' ' {
			i++
		}
		if i >= n {
			break // reached end of string
		}

		ch := s[i] // current character (byte)

		// ── Is it a number? ───────────────────────────────────
		if isDigitStart(ch) {
			// Parse the full number (may be multi-digit).
			num := 0
			for i < n && s[i] >= '0' && s[i] <= '9' {
				num = num*10 + int(s[i]-'0')
				i++
			}
			// PUSH the number onto the stack.
			stack = append(stack, num)

		// ── Is it an operator? ────────────────────────────────
		} else if ch == '+' || ch == '-' || ch == '*' || ch == '/' || ch == '%' {

			// We need at least 2 numbers on the stack to apply an operator.
			if len(stack) < 2 {
				fmt.Println("Error")
				return
			}

			// POP two values.
			// The LAST pushed is the RIGHT operand (b).
			// The second-to-last is the LEFT operand (a).
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2] // remove top two elements

			var result int

			// Apply the operator.
			// "switch" checks multiple conditions like a multi-way if.
			switch ch {
			case '+':
				result = a + b
			case '-':
				result = a - b
			case '*':
				result = a * b
			case '/':
				if b == 0 {
					fmt.Println("Error") // division by zero
					return
				}
				result = a / b
			case '%':
				if b == 0 {
					fmt.Println("Error")
					return
				}
				result = a % b
			}

			// PUSH the result back.
			stack = append(stack, result)
			i++ // move past the operator character

		} else {
			// Unknown token → invalid expression.
			fmt.Println("Error")
			return
		}
	}

	// The stack must have exactly ONE number left = the final answer.
	if len(stack) != 1 {
		fmt.Println("Error")
		return
	}

	fmt.Println(stack[0])
}

// ─── HOW TO RUN ──────────────────────────────────────────────
// go run 19_rpncalc.go "1 2 * 3 * 4 +"   → 10
// go run 19_rpncalc.go "1 2 3 4 +"        → Error  (too many operands)
// go run 19_rpncalc.go ""                  → Error
