// package main

// import "fmt"

// func main() {
// 	// question := "¿Cómo estás?"
// 	// fmt.Println(len(question))
// 	// fmt.Println(question[:6])
// 	// fmt.Println([]byte(question))
// 	// fmt.Println(string([]byte(question[:7])))

// 	// тема: Указатели
// 	// a := 10

// 	// b := a

// 	// b -= 1

// 	// fmt.Println(b, a)

// 	/*a := 10
// 	var p *int //указатель на int
// 	p = &a     // взятие адреса переменной a
// 	fmt.Println(a, p)

// 	fmt.Println(*p)
// 	*p = 15
// 	fmt.Println(a)*/

// 	// // нельзя делать указатель на константное значение и строку
// 	// const c = 5
// 	// p1 := &c
// 	// p2 := &"adidas"

// 	planets := []string{"Венера", "Земля", "Марс"} //массив - слайс
// 	planets = append(planets, "Юпитер", "Сатурн", "Уран")

// 	fmt.Println(planets)
// 	fmt.Println(len(planets)) //посмотреть длинну массива planets
// }

// func hyperspace(worlds []string) {
// 	for i := range worlds {
// 		worlds[i] = strings.TrimSpace(worlds[i])
// 	}
// }

package main

import "fmt"

func dump(label string, slice []string) {
	fmt.Printf("%v: длина %v, вместимость %v %v\n", label, len(slice), cap(slice), slice)
}
func main() {
	planets := []string{"Венера", "Земля", "Марс", "Нептун"} //[3]string{..}
	// planets = append(planets, "Юпитер", "Сатурн", "Уран")
	// //[4]string{..}
	dump("planets", planets)
	planets = append(planets, "Уран")
	dump("planets+Uran", planets)

	planetsKind := make([]string, 0, 10)
	dump("Kind Planets", planetsKind)

}
