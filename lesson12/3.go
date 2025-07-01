// Создать 5 горутин. и канал типа string. в канал записать города: москва, питер, сочи, омск, Екб.
// В каждом гофере считать из канала строку и ывести на экран: Гофер №() в городе ""

package main

import (
	"fmt"
	"time"
)

func gopherInCity(id int, cityChan chan string) {
	city := <-cityChan
	fmt.Printf("Гофер №%d в городе \"%s\"\n", id, city)
}
func main() {
	cities := []string{"Москва", "Питер", "Сочи", "Омск", "Екб"}
	cityChan := make(chan string)
	for i := 1; i <= 5; i++ {
		go gopherInCity(i, cityChan)
	}
	for _, city := range cities {
		cityChan <- city
	}
	time.Sleep(1 * time.Second)
}
