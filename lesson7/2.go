package main

import (
	"fmt"
	"math"
)

func main() {
	// int -> float
	// float -> int

	fmt.Println(math.MaxInt8)
	fmt.Println(math.MinInt8)

	myVar := 2.2                       // float64
	fmt.Println(myVar < math.MaxInt32) // float64 < maxInt32?
}
