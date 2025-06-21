package main

import (
	"fmt"
)

func main() {
	// int -> float
	// float -> int
	intVar := 360
	floatVar := 360.2
	//res := int(floatVar)
	//fmt.Println(res)
	fmt.Println(intVar <= int(floatVar))

}
