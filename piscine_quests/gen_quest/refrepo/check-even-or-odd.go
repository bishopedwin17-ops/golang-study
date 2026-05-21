package refrepo

import "fmt"

func CheckOdd() {
	// Check if a number is even or odd
	a := 7

	if a/2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
