package main

import (
	"fmt"
	"math"
)

/*
Ejercicio: Calcule e imprima el área y el perímetro del triángulo
Crear un programa que solicite al usuario que ingrese los lados de un triángulo rectángulo y luego calcule e
imprima el área y el perímetro del triángulo.

El programa debe:

Solicitar al usuario que ingrese la longitud de los dos lados del triángulo rectángulo.

Calcular la hipotenusa del triángulo usando el teorema de Pitágoras.

Calcular el área del triángulo usando la fórmula base x altura / 2.

Calcular el perímetro del triángulo sumando los lados.

Imprimir el área y el perímetro del triángulo con dos decimales de precisión.

El programa debe usar variables para almacenar los lados del triángulo, la hipotenusa, el área y el perímetro.
También debe usar constantes para representar el número de decimales de precisión que se desean en la salida.

Además, se deben utilizar funciones del paquete Math de Go para calcular la raíz cuadrada y cualquier otro cálculo
matemático necesario.
*/

const precision = 2

func main() {
	var a, b float64
	fmt.Println("Ingreso los valores para los lados \"a\" y \"b\" de su triángulo rectangulo:")
	fmt.Println("Ingrese el valor para \"a\":")
	fmt.Scanln(&a)
	fmt.Println("Ingrese el valor para \"b\":")
	fmt.Scanln(&b)
	fmt.Println(hipotenusa(a, b))
	fmt.Printf("Hipotenusa: %.*f\n", precision, hipotenusa(a, b))
	fmt.Printf("Area: %.*f\n", precision, area(a, b))
	fmt.Printf("Perimetro: %.*f\n", precision, perimetro(a, b, hipotenusa(a, b)))
}

func hipotenusa(a, b float64) float64 {
	return math.Sqrt(math.Pow(a, 2) + math.Pow(b, 2))
}

func area(a, b float64) float64 {
	return a * b / 2
}

func perimetro(a, b, c float64) float64 {
	return a + b + c
}
