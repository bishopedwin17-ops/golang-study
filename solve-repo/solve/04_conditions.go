package solve

import "os"

func CheckNumber(a int) {
	if a < 0 {
		println("number is negative")
	} else if a == 0 {
		println("number is zero")
	} else {
		println("number is positive")
	}
}

func GradeSys(score int) {
	if score < 39 {
		println("F")
	} else if score >= 40 && score < 49 {
		println("E")
	} else if score >= 50 && score < 59 {
		println("D")
	} else if score >= 60 && score < 69 {
		println("C")
	} else if score >= 70 && score < 79 {
		println("B")
	} else if score >= 80 && score <= 100 {
		println("A")
	} else {
		println("invalid score")
	}
}

func Agechecker(age int) {
	if age < 0 {
		println("invalid age")
		os.Exit(1)
	} else if age >= 0 && age <= 12 {
		println("age: %v u are underacge", "exiting", age)
		os.Exit(1)
	} else if age > 12 && age < 18 {
		println("teenager:", "come back when you're 18")
	} else if age > 18 && age <= 59 {
		println("Adult")
	} else if age >= 60 {
		println("Senior")
	}
}
