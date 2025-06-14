package main

import "fmt"

func main() {
	var a uint
	fmt.Print("введите возраст")
	fmt.Scanln(&a)
	if a < 5 {
		fmt.Println("Билет бесплатный")
	} else if a > 60 {
		fmt.Println("бесплатно")
	} else {
		fmt.Println("с вас 68 рублей")
	}
}

