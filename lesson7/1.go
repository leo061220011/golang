package main

import (
	"fmt"
	"time"
)

/*
func myFuncReturnsError() (string, error) {
	return "I am String"
}

func main() {

	err := fmt.Errorf("I am Error!")
	fmt.Println(err)

*/

func main() {

	/*fmt.Println("Hello!" + "I am Go")

	fmt.Println("Hello!" + "12")*/
	/*
		intVar := 10
		floatVar := 10.2
		fmt.Println(intVar + int(floatVar))
	*/
	/*
		intVar := 260
		floatVar := 260.2
		fmt.Println(int(floatVar) > intVar)
	*/
	/*myVar := 2.2
	fmt.Println(myVar < math.MaxInt32)
	*/

	/*myVar := 2.2
	fmt.Printf("Значение %v имеет тип %T", int(myVar), int(myVar))
	*/
	/*
		myintVar := 960
		myIntToStringVar := strconv.Itoa(myintVar)

		fmt.Println(string(myintVar)) */

	/*var a int
	fmt.Scanln(&a)

	if a >= math.MinInt8 && a <= math.MaxInt8 {
		fmt.Println("true")
	} else {
		fmt.Println("false")

	}
	*/
	/* var temperature float64
	temperature = 36.6
	fmt.Println("Температура пациента: " + fmt.Sprintf("%.1f", temperature))
	*/

	/*launch := true // можно заменить на false
	var yesNo string
	if launch {
		yesNo = "Так точно!"

	} else {
		yesNo = "Никак нет!"
	}
	fmt.Println("Ready for launch:", yesNo) */
	/*
		// работа со временем
		//  01 02 03 04 05 06
		// месяц число час(pm) мин сек 2006 года
		// 02.01.2006 15:04:05
		timeBase := "2006-01-02"
		nowDate := "2025-05-19"

		timeParced, err := time.Parse(timeBase, nowDate)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(timeParced)
	*/

	timeNow := time.Now()

	fmt.Println(timeNow.Format("02/03/06"))
	fmt.Println(timeNow.Format("Jan 02 2006 15-04"))
}
