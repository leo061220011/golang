package main

import (
	"fmt"
)

func Interation(t *int) {
	//t - локальная копия p
	//Улица Ленина дом 35
	//ул.Ленина д.35
	*t += 1
}
func main() {
	d := 3
	//var p *int
	p := &d
	Interation(p)
	fmt.Println(d) // d = 4
}
