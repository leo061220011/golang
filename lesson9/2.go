package main

import (
	"fmt"
)

func Coordinates() (int, int, int) {
	return 2, 4, 5
}

func main() {
	x, y, z := Coordinates()

	fmt.Printf("Широта: %d Долгота: %d Высота: %d\n", x, y, z)
}
