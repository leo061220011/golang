// package main

// import "fmt"

// func Interation(t int) int {
// 	t += 1
// 	return t

// 	//fmt.Println("Hi, Irina!")
// 	//fmt.Println(t)

// }

// func main() {

// 	//Cube() // Вызов функции с выводом на экран
// 	d := 3
// 	d = Interation(d)
// 	fmt.Println(d)

// }

package main

import (
	"fmt"
)

func Interation(t int) int {
	t += 1
	return t
}
func main() {
	d := 3
	d = Interation(d)
	fmt.Println(d) // d = 4
}
