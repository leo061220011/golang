/*package main

import "fmt"


func main() {
	//error type
	err := fmt.Errorf("I am error!")
	fmt.Println(err)

} */

package main

import "fmt"

func myFuncReturnsError() (string, error) {
	return "I am String!", nil
}
func main() {
	//error type

	myString, myError := myFuncReturnsError()
	if myError != nil {
		fmt.Println(myError)
	}

	fmt.Println(myString)

}
