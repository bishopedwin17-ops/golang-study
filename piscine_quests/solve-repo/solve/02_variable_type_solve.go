package solve

import "fmt"

var name string = "Gopher"
var age int = 23
var height float64 = 1.75
var isStudent bool = true

func PrintValues() {
	fmt.Printf("Name: %T %v, \n Age: %T %v, \n Height: %T %v, \n Is Student: %T %v, \n", name, name, age, age, height, height, isStudent, isStudent)
}
