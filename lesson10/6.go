// структура через композитные литералы
package main

import "fmt"

func main() {
	type location struct {
		lat  float64
		long float64
	}
	var spirit location // Повторное использование типа location
	spirit.lat = -14.5684
	spirit.long = 175.472636
	// var opportunity location // Повторное использование типа location
	// opportunity.lat = -1.9462
	// opportunity.long = 354.4734
	var opportunity location = location{lat: -1.9462, long: 354.4734}
	fmt.Println(spirit, opportunity) // Выводит: {-14.5684 175.472636} {-1.9462 354.4734}
	// var myArr [3]int
	// myArr[0] = 2
}
