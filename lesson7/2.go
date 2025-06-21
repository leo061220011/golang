package main

import (
	"fmt"
)

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Div by 0!")
	}
	return a / b, nil

}

func main() {

	a, b := 2, 1
	result, err := div(a, b)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)

}
