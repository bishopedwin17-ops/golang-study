package refrepo

import "fmt"

func LoopAtoz() {
	// Use a for loop to print alphabets from a to z
	for i := 'a'; i <= 'z'; i++ {
		fmt.Printf("%c", i)
	}
	fmt.Println('\n')
}
