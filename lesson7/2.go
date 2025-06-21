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
	return "I am String", fmt.Errorf("I am error!")
}
func main() {
	//error type

	myString, myError := myFuncReturnsError()
	fmt.Println(myString, " ", myError)

}
