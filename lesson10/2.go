// ## Задача 2: Циклы
// Создайте Слайс! и выведете на экран все элементы в обратном порядке

package main

import (
	"fmt"
)

func main() {

	//вариант Павла
	// myFirstArr := []int{1, 2, 3, 4}

	// fmt.Println("Элементы в обартном порядке:")
	// for i := len(myFirstArr) - 1; i >= 0; i-- {
	// 	fmt.Println(myFirstArr[i])
	// }

	//мой вариант

	myArr := []string{"мама", "папа", "я", "сестра"}

	fmt.Println("Элементы в обартном порядке:")
	for i := len(myArr) - 1; i >= 0; i-- {
		fmt.Println(myArr[i])
	}
}
