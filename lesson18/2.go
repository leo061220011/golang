// Функция, принимает 2 параметра типа int и возвращает результат деления 1 на 2
// В ДЗ вы обрабатывали ошибку деления на 0
// Теперь давайте обработаем деление на 0 с помощью паники.
// Вызовите панику при b = 0 и отловите ее
// func div(a,b int) float64 {
// 	if b == 0{
// 		panic("div by zero") // и обработать панику!
// 	}
// }
// В main:
// 	div(1,0) // Я обработал панику: div by zero

package main

import "fmt"

func div(a, b int) float64 {
	if b == 0 {
		panic("div by zero")
	}
	return float64(a) / float64(b)
}

func main() {

	defer func() {
		if d := recover(); d != nil {
			fmt.Printf("Я обработал панику: %v\n", d)
		}
	}()

	fmt.Println(div(1, 0))
}
