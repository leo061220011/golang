package main

import "fmt"

func main() {

	// && - И
	// || ИЛИ
	// ! НЕ
	a := 3
	b := 2
	bfunc := func() bool {
		b = b + 1
		return true
	}

	if a == 1 || bfunc() {
		fmt.Println("Yes")
		fmt.Println(b)
	}

}
