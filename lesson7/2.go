package main

import (
	"fmt"
	"strconv"
)

func main() {

	myIntVar := "960"

	myIntToStringVar, err := strconv.Atoi(myIntVar)

	if err != nil {
		fmt.Println("Я не могу это перевести!")
		return
	}

	fmt.Printf("%v - %T", myIntToStringVar, myIntToStringVar)

}
