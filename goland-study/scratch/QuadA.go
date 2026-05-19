package main

import (
	"fmt"
	"strings"
)

//define the function QuadA, write if for x and y when less that 1
func QuadA(x int, y int) {
	if x < 1 || y < 1 {
	return
	}
	// write a variable midwith and write the value of the midwith to be 2 less than the value of x
	midwith := x - 2
	// if the midwith is anything less than 0 then the midwith should return 0 meaning the middle line is empty
	if midwith < 0 {
		midwith = 0
	}
	//write the variable of the top line/row and define its function
	//also repeat the middle string as many times as needed 
	top := "o" + strings.Repeat("-", midwith)
	if x > 1 {
		top += "o"
	}

	mid := "|" + strings.Repeat(" ", midwith)
	if x > 1 {
		top += "|"
	}

	fmt.Println(top)
	// r = row
	// r := 0 start counting the loop r from the value of 0 and repeat
	// r < y-2 so far the value of y - 2 is less than the value of r rerun the loop again till the value of r is equal to the value of y - 2
	// r++ add 1 to the value of r after every successful loop
	for r := 0; r < y - 2; r++ {
		fmt.Println(mid)
	}
	if y > 1 {
	fmt.Println(top)
	}
}

func main() {
	QuadA(1, 5)
}