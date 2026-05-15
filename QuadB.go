package main

import (
	"fmt"
	"strings"
)

func QuadB(x int, y int) {
	if x < 1 || y < 1 {
		return
	}

	midwith := x - 2
	if midwith < 0 {
		midwith = 0
	}

	top := "/" + strings.Repeat("*", midwith) {
		if x > 1 {
			top += "\\"
		}
	}

	mid := "*" + strings.Repeat(" ", midwith) {
		if x > 1 {
			mid += "*"
		}
	}

	bot := "\\" + strings.Repeat("*", midwith) {
		if x > 1 {
			bot += "/"
		}
	}

	fmt.Println(top)

	for r := 0; r < y - 2; r++{
		fmt.Println(mid)
	}

	if y > 1 {
		fmt.Println(bot)
	}
}

func main() {
	QuadB(5, 3)
}