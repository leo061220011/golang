package main

import (
	"fmt"
	"strings"
)

type Cat struct {
	name  string
	voice string
}
type Dog struct {
	name  string
	voice string
}
type Bird struct {
	name  string
	voice string //курлык!
}

func (c Cat) speak() string {
	return c.voice
}
func (d Dog) speak() string {
	return d.voice
}
func (b Bird) speak() string {
	return "Kesha-good"
}
func (b Bird) kurlick() string {
	return b.voice
}

type Speaker interface {
	speak() string
	kurlick() string
}

func SpeakTwice(s Speaker) {
	fmt.Println(strings.Repeat(s.speak(), 2))
}
func main() {
	barsik := Cat{"Барсик", "Мяу!"}
	sharik := Dog{"Шарик", "Гав!"}
	Kesha := Bird{"Кеша", "Kurlick!"}
	SpeakTwice(barsik)
	SpeakTwice(sharik)
	SpeakTwice(Kesha)
}
