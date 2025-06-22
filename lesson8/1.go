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

package main

import (
	"fmt"
)

func main() {
	planets := [8]string{
		"Меркурий",
		"Венера",
		"Земля",
		"Юпитер",
		"Сатурн",
		"Уран",
		"Нептун",
	}

	fmt.Println(planets)

}
