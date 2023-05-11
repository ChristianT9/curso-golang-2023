package main

import "fmt"

func main() {
	a := 20
	b := 10

	//Suma
	result := a + b
	fmt.Println("Suma:", result)

	//Resta
	result = a - b
	fmt.Println("Resta:", result)

	//Multiplicación
	result = a * b
	fmt.Println("Multiplicación:", result)

	//División
	result = a / b
	fmt.Println("División:", result)

	div := 3.0 / 2.0
	fmt.Println("División:", div)

	//Modulo
	result = a % b
	fmt.Println("Modulo:", result)

	result = 3 % 2
	fmt.Println("Modulo:", result)
}
