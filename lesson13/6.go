package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func pr() {
	fmt.Println("Hi, Boris")
	wg.Done()
}
func main() {
	wg.Add(1)
	go pr()
	wg.Wait()
}

// https://github.com/Go-cource/lesson12
