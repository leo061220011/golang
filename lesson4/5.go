package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {

		if i == 0 {
			fmt.Printf("%d - Ноль\n", i)
			continue // пропускает одну итерацию
		}

		if i%2 == 0 {

			fmt.Printf("%d - четное\n", i)

		} else {
			fmt.Printf("%d - нечетное\n", i)
		}
	}

}
