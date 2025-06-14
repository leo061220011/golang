package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	x := 5
	y := 3
	result := add(x, y)
	fmt.Println("Сумма =", result)

}
