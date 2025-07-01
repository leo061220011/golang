// package main

// import (
// 	"fmt"
// 	"time"
// )

// func sleepyGopher(i int, c chan int) {
// 	time.Sleep(3 * time.Second)
// 	fmt.Println("... gooo ... i am gopher #", i)
// 	c <- i
// }
// func main() {
// 	c := make(chan int)
// 	for i := 0; i < 5; i++ {
// 		go sleepyGopher(i, c)
// 	}
// 	for i := 0; i < 5; i++ {
// 		gopherId := <-c
// 		fmt.Println("gopher # ", gopherId, "from main")
// 	}
// }

// Создать 5 горутин. и канал типа string. в канал записать города: москва, питер, сочи, омск, Екб.
// В каждом гофере считать из канала строку и ывести на экран: Гофер №() в городе ""

package main

import (
	"fmt"
)

func gopherInCity(id int, cityChan chan string, intChan chan int) {
	city := <-cityChan
	fmt.Printf("Гофер №%d в городе \"%s\"\n", id, city)
	intChan <- id
}
func main() {
	cities := []string{"Москва", "Питер", "Сочи", "Омск", "Екб"}
	cityChan := make(chan string)
	intChan := make(chan int)
	for i := 1; i <= 5; i++ {
		go gopherInCity(i, cityChan, intChan)
	}
	for _, city := range cities {
		cityChan <- city
		<-intChan
	}
}
