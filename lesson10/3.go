// ## Задача 3: Функции
// Создайте функции, которые переводят градусы Цельсия в градусы Кельвин  (Целс. 0 = Кельв. 273,15) и обратно. Выведете на экран:
// %v градусов Цельсия - %v градусов Кельвин

package main

import "fmt"

func cels_kelv(c float64) float64 {
	return c + 273.15
}
func kelv_cels(k float64) float64 {
	return k - 273.15
}

func main() {

	celsiy := 36.6
	kelvin := 200.0

	fmt.Printf("%v Cels - %v Kel\n", celsiy, cels_kelv(celsiy))

	fmt.Printf("%v Kelv - %.2f Cels\n", kelvin, kelv_cels(kelvin))

}
