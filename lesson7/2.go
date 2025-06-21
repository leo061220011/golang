/*package main

import "fmt"


func main() {
	//error type
	err := fmt.Errorf("I am error!")
	fmt.Println(err)

} */

package main

import (
	"errors"
	"fmt"
)

func myFuncReturnsError() (string, error) {
	return "I am String!", errors.New("I am errors error")
}
func main() {
	//error type

	myString, _ := myFuncReturnsError() //забили на ошибку
	/*	if myError != nil {
		fmt.Println(myError)
	} */

	fmt.Println(myString)

}
