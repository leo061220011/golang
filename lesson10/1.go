// Сохраните в переменную настоящее время и отформатируйте по шаблону:
// 1. 2025/06/26 18:36:00
// 2. 26-06-26
// 3. 26 Jun 2025 18:37

// now := time.Now()

// 	format1 := now.Format("2006/01/02 15:04:05")
// 	format2 := now.Format("02/01/06")
// 	format3 := now.Format("02 Jan 2006 15:04")

// 	fmt.Println("1.", format1)
// 	fmt.Println("2.", format2)
// 	fmt.Println("3.", format3)

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {

// timeBase := "2025-06-26"
// nowDate := "2025-06-26"

// timeParced, err := time.Parse(timeBase, nowDate)
// if err != nil {
// 	fmt.Println(err)
// }
// fmt.Println(timeParced)

// timeNow := time.Now()

// fmt.Println(timeNow.Format("02/03/06 15-04-05"))
// fmt.Println(timeNow.Format("02-03-06"))
// fmt.Println(timeNow.Format("Jan 02 2006 15-04"))

// вариант учителя
package main

import (
	"fmt"
	"time"
)

func main() {

	timeNow := time.Now()
	fmt.Println(timeNow.Format("2006/01/02 15:04:05"))
	fmt.Println(timeNow.Format("02-01-06"))
	fmt.Println(timeNow.Format("02 Jan 2006 15:04"))
}
