package main

import "fmt"

func main() {
	var name, name2 string
	var age int

	fmt.Print("Otro mensaje") // No inserta un salto de linea

	fmt.Print("\nIngrese su nombre: ")
	fmt.Scanln(&name, &name2)
	fmt.Print("Ingrese su edad: ")
	fmt.Scanln(&age)

	fmt.Printf("\nHola, me llamo %s %s y tengo %d años.\n", name, name2, age) // Da formato al mensaje

	greeting := fmt.Sprintf("Hola, me llamo %s y tengo %d años.", name, age)
	fmt.Println(greeting)

	fmt.Printf("El tipo de name es: %T\n", name)
	fmt.Printf("El tipo de age es: %T\n", age)
}
