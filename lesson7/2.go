package main

import (
	"fmt"
)

func main() {
	// int -> float
	// float -> int
	intVar := 10
	floatVar := 10.2

	fmt.Println(intVar + int(floatVar))

}
