// type celsius float64
// const degrees = 20
// var temperature celsius = degrees
// temperature += 10
// fmt.Println(temperature)
// var floatVar float64 = 10.0
// temperature += celsius(floatVar)
// fmt.Printf("%T", temperature)

package main

import "fmt"

type celsius float64
type kelvin float64

// kelvinToCelsius converts °K to °C
func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15) // Необходима конвертация типа
}
func main() {
	var k kelvin = 294.0 // Аргумент должен быть типа kelvin
	c := kelvinToCelsius(k)
	fmt.Print(k, "° K is ", c, "° C") // Выводит: 294° K is 20.850000000000023° C
	var cel celsius = 36.0
}
