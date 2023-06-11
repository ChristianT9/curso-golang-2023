package main

import (
	"fmt"
	"math"
)

func main() {
	a := 10
	b := 3

	fmt.Println("Operaciones básicas")
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(float64(a) / float64(b))
	fmt.Println(a % b)
	fmt.Println()

	b++ // Incremento
	fmt.Println("Incremento de b")
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(float64(a) / float64(b))
	fmt.Println(a % b)
	fmt.Println()

	b-- //Decremento
	fmt.Println("Decremento de b")
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(float64(a) / float64(b))
	fmt.Println(a % b)
	fmt.Println()

	a = a + 5
	fmt.Println("Sumando un valor al valor base de a")
	fmt.Println(a)
	a += 5 // Suma en asignación. También se puede usar (-=, *=, %=, /=)
	fmt.Println(a)
	fmt.Println()

	fmt.Println("Utilizando el paquete math")
	fmt.Println(math.Pi)
	fmt.Println(math.E)
	fmt.Println(math.Pow(2, 4))
	fmt.Println(math.Sqrt(64))
	fmt.Println(math.Cbrt(27))
}
