package main

import (
	"fmt"
	"strconv"
)

func main() {
	var integer16 int16 = 50
	var integer32 int32 = 100

	fmt.Println(int32(integer16) + integer32)

	s := "100"
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Falló la conversión de string a int")
	} else {
		fmt.Println(n + n)
	}

	i := 42
	s = strconv.Itoa(i)
	fmt.Println(s + s)
}
