package main

import (
	"fmt"
	"strings"
)

// пишем конвейер:
// 1 го рутина - генератор. Она генерирует слова и передает их в канал для анализа
// 2 горутина считывает слова и делает анализ. В случае, если найдено слово-конец генерации- останавливаемся
// 3 - выводит слова на экран
func sourceGopher(ch1 chan string) {
	for _, v := range []string{"hello world", "a bad apple", "goodbye all", "bad day =(", "my friends"} {
		ch1 <- v
	}
	ch1 <- ""
}
func filterGopher(ch1, ch2 chan string) {
	for {
		item := <-ch1
		if item == "" {
			ch2 <- ""
			return
		}
		if !strings.Contains(item, "bad") {
			ch2 <- item
		}
	}
}
func printGopher(ch2 chan string) {
	for {
		v := <-ch2
		if v == "" {
			return
		}
		fmt.Println(v)
	}
}
func main() {
	/*
		generate -> (строки) -> analyze -> (строки после анализа) -> print
	*/
	ch1 := make(chan string)
	ch2 := make(chan string)
	go sourceGopher(ch1)
	go filterGopher(ch1, ch2)
	printGopher(ch2)
}
