/*package main

import "fmt"

func main() {
	var planets [8]string
	planets[0] = "Меркурий"
	planets[1] = "Венера"
	planets[2] = "Земля"

	earth := planets[2]
	fmt.Println(planets)
	fmt.Println("Наша планета: ", earth)

	fmt.Println(len(planets))
	fmt.Println(planets[5] == "")
}
*/
/*
package main

import (
	"fmt"
)

func main() {
	planets := [9]string{
		"Меркурий",
		"Венера",
		"Земля",
		"Юпитер",
		"Сатурн",
		"Уран",
		"Нептун",
		"центавра",
	}

	fmt.Println(planets)
	for i := 0; i < len(planets); i++ {
		fmt.Println(planets[i])
	}

}
*/

package main

import (
	"fmt"
)

func main() {
	taskArray := [5]int{
		82, 97, 110, 103, 101,
	}
	for _, r := range taskArray {
		fmt.Printf("%c", r)
	}

}
