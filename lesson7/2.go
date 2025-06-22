package main

import (
	"fmt"
	"strconv"
)

func main() {

	myIntVar := 960
	myIntToStringVar := strconv.Itoa(myIntVar)
	fmt.Println("I have " + myIntToStringVar + " Dollars")

}
