// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	v := Abs(3)
// 	fmt.Println(v)
// }

// // Abs возвращает абсолютное значение.
// // Например: 3.1 => 3.1, -3.14 => 3.14, -0 => 0.
// func Abs(value float64) float64 {
// 	return math.Abs(value)
// }

// // go get github.com/stretchr/testify команда для создания пакета тестифай

package main

import (
	"fmt"
	"math"
)

func main() {
	v := Abs(3)
	fmt.Println(v)
}

// Abs возвращает абсолютное значение.
// Например: 3.1 => 3.1, -3.14 => 3.14, -0 => 0.
func Abs(value float64) float64 {
	return math.Abs(value)
}
func RetTrue() bool {
	return true
}
