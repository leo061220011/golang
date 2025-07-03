package main

import (
	"fmt"
	"sync"
	"time"
)

// мьютексы
var myMap = map[string]int{"qwe": 1, "ewq": 2}
var mu sync.Mutex

func gor1() {
	for i := 0; i < 1000; i++ {
		mu.Lock()
		myMap["qwe"] = i
		mu.Unlock()
	}
}
func gor2() {
	for i := 0; i < 1000; i++ {
		mu.Lock()
		myMap["ewq"] = i
		mu.Unlock()
	}
}
func main() {
	go gor1()
	go gor2()
	fmt.Println(myMap)
	time.Sleep(1 * time.Second)
}
