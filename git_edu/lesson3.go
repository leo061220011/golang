package main

import (
	"fmt"
)

func main() {
	var name string
	var age int
	fmt.Print("Введите ваше имя: ")
	fmt.Scanln(&name)
	fmt.Print("Введите ваш возраст: ")
	fmt.Scanln(&age)
	fmt.Printf("Привет, %s! Вам %d лет.\n", name, age)
}





%s - строка.
	%d - целое число.
	%f - число с плавающей точкой.
	%t - булево значение.
	%v - значение в стандартном формате (универсальный спецификатор).