package solve

import "fmt"

func AskInput() {
	var firstname string
	fmt.Println("input first name")
	_, err := fmt.Scan(&firstname)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	// var lastname string
	// fmt.Println("input last name")
	// _, err := fmt.Scan(&lastname)
	// if err != nil {
	// 	fmt.Println("Error reading input:", err)
	// 	return
	// }
	// fmt.Println("Hello,", firstname, lastname)

	var age int
	fmt.Println("input age")
	_, err = fmt.Scan(&age)
	if err != nil {
		fmt.Println("Invalid age input:", err)
		return
	}
	fmt.Printf("Hello, %v, you are %v years old\n", firstname, age)
}

func SimpleCalculator() {
	var num1, num2 int
	var operator string

	fmt.Print("Enter first number: ")
	if _, err := fmt.Scan(&num1); err != nil {
		fmt.Println("Invalid number:", err)
		return
	}

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scan(&operator)

	fmt.Print("Enter second number: ")
	if _, err := fmt.Scan(&num2); err != nil {
		fmt.Println("Invalid number:", err)
		return
	}

	// Perform the math
	switch operator {
	case "+":
		fmt.Print("Result: ", num1+num2)
	case "-":
		fmt.Print("Result: ", num1-num2)
	case "*":
		fmt.Print("Result: ", num1*num2)
	case "/":
		if num2 == 0 || num1 == 0 {
			fmt.Println("Cannot divide by zero")
		} else {
			fmt.Printf("Result: %d\n", num1/num2)
		}
	default:
		fmt.Println("Unknown operator")
	}
}
