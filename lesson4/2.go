package main

import "fmt"

func main() {

	v := 0
	//for i := 1; i < 10; i++ {
	for {
		v++
		fmt.Println(v)

		if v == 7 {

			break
		}

		//break -  принудительная остановка цикла
	}
}
