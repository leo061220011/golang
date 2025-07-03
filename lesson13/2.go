package main

import (
	"fmt"
	"strings"
)

// пишем конвейер:
// 1 го рутина - генератор. Она генерирует слова и передает их в канал для анализа
// 2 горутина считывает слова и делает анализ. В случае, если найдено слово-конец генерации- останавливаемся
// 3 - выводит слова на экран
func sourceGopher(downstream chan string) {
	for _, v := range []string{"hello world", "a bad apple", "goodbye all", "bad day =(", "my friends"} {
		downstream <- v
	}
	downstream <- ""
}
func filterGopher(upstream, downstream chan string) {
	for {
		item := <-upstream
		if item == "" {
			downstream <- ""
			return
		}
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
}
func printGopher(upstream chan string) {
	for {
		v := <-upstream
		if v == "" {
			return
		}
		fmt.Println(v)
	}
}
func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go sourceGopher(ch1)
	go filterGopher(ch1, ch2)
	printGopher(ch2)
}
