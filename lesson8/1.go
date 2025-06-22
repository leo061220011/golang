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

/*package main

import (
	"fmt"
)

func main() {
	taskArray := [...]int{82, 97, 110, 103, 101}
	for _, v := range taskArray {
		fmt.Print(string(rune(v)))
	}

}*/

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

	giants := planets[4:8]
	gas := giants[0:2]
	ice := giants[2:4]

	fmt.Println(giants, gas, ice)
	ice[0] = "Нептун"
	fmt.Println(planets)

}
