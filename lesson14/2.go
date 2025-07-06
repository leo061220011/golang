package main

import (
	"fmt"
	"os"
)

func main() {

	data, err := os.ReadFile("example_txt.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)

}

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	data, err := os.ReadFile("example_txt.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(string(data))
// }
