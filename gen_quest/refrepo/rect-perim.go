package refrepo

import "fmt"

func RectPerimeter(a, b int) int{
	warn := " - Error: can't have a negative perimeter value"
	if a < 0 {
		fmt.Print("width: -1",warn, "\n")
	} else {
		fmt.Print("width: ", a, "\n")
	}
	
	if b < 0 {
		fmt.Print("width: -1", warn, "\n")
	} else {
		fmt.Print("width: ", b, "\n")
	}
	return a
}