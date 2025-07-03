// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// func sleepyGopher(i int, c chan int) {
// 	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
// 	c <- i
// }
// func main() {
// 	c := make(chan int)
// 	for i := 0; i < 5; i++ {
// 		go sleepyGopher(i, c)
// 	}
// 	timeout := time.After(3 * time.Second)
// 	// for i := 0; i < 5; i++ {
// 	// 	gopherId := <-c
// 	// 	fmt.Println("gopher # ", gopherId, "from main")
// 	// }
// 	select {
// 	case gopherId := <-c:
// 		fmt.Println("gopher # ", gopherId, "from main")
// 	case <-timeout:
// 		fmt.Println("time is out")
// 	}
// }

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sleepyGopher(i int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	c <- i
}
func main() {
	c := make(chan int)
	for i := 0; i < 5; i++ {
		go sleepyGopher(i, c)
	}
	timeout := time.After(2 * time.Second)
	for i := 0; i < 5; i++ {
		select {
		case gopherID := <-c:
			fmt.Println("gopher #", gopherID, "")
		case <-timeout:
			fmt.Println("time is out")
			return
		}
	}
}
