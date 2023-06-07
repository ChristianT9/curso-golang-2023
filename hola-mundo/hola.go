package main

import (
	"fmt"
	"rsc.io/quote"
)

// Declaración de constantes
const (
	Pi  float32 = 3.14
	Pi2         = 3.14
	x           = 100
	y           = 0b1010 //Binario
	z           = 0o12   //Octal
	w           = 0xFF   //Hexadecimal
)

const (
	Doming = iota + 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
)

/*var (
	firstName string = "Christian Alberto"
	lastName  string = "Tamayo Robayo"
	age       int    = 23
)*/

func main() {
	fmt.Println("Hola Mundo")
	fmt.Println(quote.Hello())
	fmt.Println(quote.Go())

	//Declaración de variables

	/*var firstName, lastName string
	var age int*/

	/*var (
		firstName, lastName string
		age                 int
	)

	firstName = "Christian Alberto"
	lastName = "Tamayo Robayo"
	age = 23*/

	/*var (
		firstName string = "Christian Alberto"
		lastName  string = "Tamayo Robayo"
		age       int    = 23
	)*/

	/*var (
		firstName = "Christian Alberto"
		lastName  = "Tamayo Robayo"
		age       = 23
	)*/

	//var firstName, lastName, age = "Christian Alberto", "Tamayo Robayo", 23

	//firstName, lastName, age := "Christian Alberto", "Tamayo Robayo", 23

	firstName := "Christian Alberto"
	lastName := "Tamayo Robayo"
	age := 23
	age = 30

	fmt.Println(firstName, lastName, age)

	fmt.Println(Pi, Pi2)
	fmt.Println(x, y, z, w)

	fmt.Println(Viernes)
}
