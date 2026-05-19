// Create a Go program to demonstrate the use of the if statement by checking if a variable x = 20 is greater than 10.
// Required Input:

// A predefined variable x = 20.

// Expected Output:

// x is greater than 10.

package refrepo

import "fmt"

func IfElse() {
	// Use an if statement to check the condition
	x := 20
	if x < 10 {
		fmt.Println("x is less than 10.")
	} else {
		fmt.Println("x is greater than 10.")
	}
}
