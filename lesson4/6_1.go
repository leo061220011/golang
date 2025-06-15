package main // Запросить 2 числа, 1>2. Считать их и вывести количество четных между ними Необязательно: вывести эти числа
import "fmt"

func main() {

	var x, y int
	z := 0

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&x)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&y)

	if x <= y {
		fmt.Println("Ошибка: первое число должно быть больше второго!")
		return
	}

	for i := x - 1; i > y; i-- {

		if i%2 != 0 {
			fmt.Printf("%d - нечетное\n", i)
			z++

		}
	}

	fmt.Printf("%d - количество нечетных", z)
}
