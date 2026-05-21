package solve

import "fmt"

func CelsiusToFahrenheit(celsius float64) float64 {
	a := celsius*9/5 + 32
	fmt.Println(a)

	return a
}

func FahrenheitToCelsius(fah float64) float64 {
	b := (fah - 32) * 5 / 9
	fmt.Println(b, "C")
	return b
}

func CelsiusToKelvin(celsius float64) float64 {
	z := celsius + 273
	fmt.Println(z, "k")
	return z
}
func KelvinToCelcius(kelvin float64) float64 {
	k := kelvin - 273
	fmt.Println(k, "C")
	return k
}

func FahrenheitToKelvin(f float64) float64 {
	fk := ((f - 32) * 5 / 9) + 273
	fmt.Printf("%.2f", fk)
	fmt.Print("K", "\n")
	return fk
}

func SwapInts(a, b int) (int, int) {
	a, b = b, a
	fmt.Println(a, b)
	return a, b
}

func SwapStrings(i, j string) (string, string) {
	i, j = j, i
	fmt.Println(i, j)
	return i, j
}
