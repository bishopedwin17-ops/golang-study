package solve

import "fmt"

func BmiCalculator() {
	var height float64
	var weight float64

	fmt.Println("input your correct height in meters")
	fmt.Println("input your correct weight in kg")
	_, err := fmt.Scan(&height, &weight)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	squareheight := height * height
	BMI := weight / squareheight
	BMI2f := fmt.Sprintf("%.3f", BMI)

	fmt.Println("your total body fat is, BMI: ", BMI2f, "BMI")
}