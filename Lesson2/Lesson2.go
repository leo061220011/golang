package main

import (
	"fmt"
	adding "lesson2/adding"
)

func Add(param1 int, param2 int) int {
	return param1 + param2
}
func main() {
	//a, b := 4, 5
	//fmt.Println(a)
	//fmt.Println(b)
	//myArray := [4]int{1,2,3,4}обычный массив
	//mySlice := []int{1,2,3,4,5}//динамический массив

	//myMap := map[int]string{1: "value1", 2: "value2"}//
	//myName := "Irina"
	//fmt.Println(myName)
	a := 1
	b := 2
	//sum := add(a, b)
	//fmt.Println(sum)
	fmt.Println(adding.Add(a, b))

}
