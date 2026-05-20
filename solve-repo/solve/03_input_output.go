package solve

import "fmt"

func AskInput() {
	var firstname string
	println("input first name")
	_, prntname := fmt.Scanln(&firstname)
	if prntname == nil {
		fmt.Println("wrong or empty input:", prntname)
		return
	}
	fmt.Println("Hello,", firstname)

	var age int
	println("input age")
	_, prntage := fmt.Scanln(&age)
	if prntage == nil {
		fmt.Println("wrong or empty input:", prntage)
		return
	}
	fmt.Println("Hello,", firstname, "you are", age, "years old")
}
