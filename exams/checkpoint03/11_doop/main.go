// =============================================================
// EXERCISE: doop
//
// QUESTION: Write a program that takes three arguments: 
//             1. A value (number)
//             2. An operator (+, -, /, *, %)
//             3. Another value (number)
//           It should print the result.
//
// CONSTRAINTS:
//           - MUST handle division by zero (print nothing or "0", 
//             check exam rules, usually "No division by 0\n").
//           - MUST handle modulo by zero (usually "No modulo by 0\n").
//           - MUST handle integer overflows. If a number doesn't 
//             fit in a standard 64-bit int, do not panic.
//           - You CANNOT use `strconv` or `fmt` (often restricted).
//             (If `fmt` is allowed, use it. But typically you must 
//              write your own `atoi` and `itoa`).
//
// EXAMPLE:
//           go run main.go 1 + 1      -> 2
//           go run main.go 10 / 0     -> No division by 0
//           go run main.go 1 +        -> (nothing, invalid args)
// =============================================================

package main

import (
	"os"
)

func main() {
	// ---------------------------------------------------------
	// TODO: 
	// 1. Check `len(os.Args)`. It must be exactly 4 
	//    (program name + 3 args).
	// 2. Parse os.Args[1] and os.Args[3] using your own Atoi().
	//    Ensure they are valid numbers.
	// 3. Switch on os.Args[2] to determine the operator.
	// 4. Perform the calculation. Catch divide/modulo by zero!
	// 5. Print the result using your own Itoa() and os.Stdout.
	// ---------------------------------------------------------
}
