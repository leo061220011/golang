package main

import (
	"fmt"
)

func main() {
	// int -> float
	// float -> int

	myVar := 2.2
	var myVarInt int = 12
	var myVarInt64 int64 = 13
	fmt.Printf("%T \n", myVar)
	fmt.Printf("%T \n", int(myVar))
	fmt.Printf("Значение %v имеет тип %T", int(myVar), int(myVar))
	fmt.Fprintln(myVarInt + myVarInt64)
}
