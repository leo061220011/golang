// создать 2 структуры: кошка, собака name:, voice.
// cоздать экземпляры этих структур
// создать методы speak() которые выводят на экран, что говорит животное
// barsik.speak() // мяу

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

func (d Dog) speak() {
	fmt.Printf("Собака %v говорит %v!\n", d.name, d.voice)

}

func (c Cat) speak() {
	fmt.Printf("Кошка %v говорит %v!\n", c.name, c.voice)
}

func speakTwiceCat(c Cat) {
	fmt.Println(strings.Repeat(c.voice+"\n", 2))
}

func speakTwiceDog(d Dog) {
	fmt.Println(strings.Repeat(d.voice+"\n", 2))
}

func main() {
	Sara := Cat{"Сара", "Мяу "}
	Ronni := Dog{"Ронни", "Гав "}
	Sara.speak()
	Ronni.speak()

	speakTwiceCat(Sara)

	speakTwiceDog(Ronni)

}
