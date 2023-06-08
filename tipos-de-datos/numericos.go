package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.MinInt64, math.MaxInt64)
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)
	fmt.Println(math.MaxUint8)
	var r rune = '♡'
	fmt.Println(r)
}
