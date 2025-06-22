package main

import (
	"fmt"
	"strconv"
)

func main() {

	myIntVar := "960hello123"

	myIntToStringVar, err := strconv.Atoi(myIntVar)

	if err != nil {
		fmt.Println("Я не могу это перевести!")
		return
	}

	fmt.Println("%v - %t", myIntToStringVar, myIntToStringVar)

}
