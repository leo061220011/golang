/*package main
import "fmt"

func main() {

	//const pi = 3.1415
	var maVar int = 5
	const myConst = myVar
}

package main
import (
	"fmt"
)
func main() {
	manyValues := 1.0 / 3
	fmt.Println(manyValues)
	fmt.Printf("%v\n", manyValues) //0.33333333333
	fmt.Printf("%f\n", manyValues)
	fmt.Printf("%.3f\n", manyValues)
	fmt.Printf("%4.2f\n", manyValues) //
	fmt.Printf("%05.2f\n", manyValues)
}



%s - строка.
%d - целое число.
%f - число с плавающей точкой.
%t - булево значение.
%v - значение в стандартном формате (универсальный спецификатор).


package main

import (
	"fmt"
)

func main() {

	a := 55.7673 / 2

	fmt.Println(a)
	fmt.Printf("%v\n", a)
	fmt.Printf("%f\n", a)
	fmt.Printf("%.3f\n", a)
	fmt.Printf("%4.2f\n", a)
	fmt.Printf("%008.2f\n", a)

}
*/

/*
package main

import (

	"fmt"
	"math"

)

	func main() {
		bank := 0.1
		bank += 0.2

		//fmt.Println(bank)
		//fmt.Printf("%.55f\n", 0.1)
		//fmt.Println(bank == 0.3) для golang 0,1 + 0,2 != 0,3
		fmt.Println(math.Abs(bank-0.3) < 0.0001)

}
*/
package main

import (
	"fmt"
	"math"
)

func main() { //Целые числа

	/*var a int8 = 127
	a++
	fmt.Println(a)*/

	/*glass := 500
	fmt.Printf("Type of %v - %T", glass, glass)*/
	/*
		fmt.Println(math.Pi)
		fmt.Println(math.E)
		fmt.Println(math.Phi)
		fmt.Println(math.Sqrt2)
		fmt.Println(math.Ln2)
		fmt.Println(math.MaxFloat64)
		fmt.Println(math.SmallestNonzeroFloat64)*/

	/*var x, y float64 = 1, 2

	fmt.Println(math.Abs(x))
	fmt.Println(math.Max(x, y))
	fmt.Println(math.Min(x, y))
	fmt.Println(math.Mod(x, y)) //остаток от деления
	fmt.Println(math.Pow(x, y)) //введение в степень
	fmt.Println(math.Sqrt(x))   //квадратный корень
	*/

	/*fmt.Printf("Модуль -5: %.1f\n", math.Abs(-5))
	fmt.Printf("Максимум из 3 и 7: %.1f\n", math.Max(3, 7))
	fmt.Printf("2 в степени 3: %.1f\n", math.Pow(2, 3))
	fmt.Printf("Квадратный корень из 16: %.1f\n", math.Sqrt(16))*/

	/*//округления
	x := 2.2345678
	fmt.Println(math.Ceil(x))  //округление до целого вверх
	fmt.Println(math.Floor(x)) //окургелние до ближайшего вниз
	fmt.Println(math.Trunc(x)) //отбросить целую часть*/

	/* практическое задание №1

	x1 := (-b + math.Sqrt(d)) / (2 * a)
	x2 := (-b - math.Sqrt(d)) / (2 * a)*/

	//задача 2) Получить от пользователя основание и степень (два числа float64). Вывести результат возведения первого в степень второго

	/*var a, b float64

	fmt.Scan(&a)
	fmt.Scan(&b)
	r := math.Pow(a, b)
	fmt.Printf("%v", r)*/

	//задача 3) на вход программа получает 3 числа. с использованием пакета math выведете на экран максимальное из них
	var a, b, c float64
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	r := math.Max(math.Max(a, b), c)
	fmt.Printf("Максимальное число: %v", r)

}
