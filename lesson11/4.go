package main

import (
	"fmt"
)

/*
создать 2 структуры: кошка, собака name:, voice.
cоздать экземпляры этих структур
создать методы speak() которые выводят на экран, что говорит животное
barsik.speak() // мяу
*/
type Cat struct {
	name  string
	voice string
}
type Dog struct {
	name  string
	voice string
}

func (d Dog) speak() {
	fmt.Printf(d.voice + "\n")
}
func (c Cat) speak() {
	fmt.Printf(c.voice + "\n")
}

/*
Расширение
	напсиать функию(функции), SpeakTwice(){
		fmt.Println(strings.Repeat("", 2))
	}
*/
// func SpeakTwiceCat(c Cat) {
// 	fmt.Println(strings.Repeat(c.voice, 2))
// }
// func SpeakTwiceDog(c Dog) {
// 	fmt.Println(strings.Repeat(c.voice, 2))
// }
func main() {
	barsik := Cat{"Барсик", "Мяу!"}
	Sharik := Dog{"Шарик", "Гав!"}
	barsik.speak()
	Sharik.speak()
	
	var animal interface {
		speak()
	}

	animal = barsik
	animal.speak()
	animal = Sharik
	animal.speak()
}
