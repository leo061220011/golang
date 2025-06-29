// 2 структуры: машина, самолет
// у каждого метод move(). Самолет - Самолет летит
// 						Машина - машина едет
// функцию принимает интерфейс, выводит "Самолет летит"+" далеко"

package main

import (
	"fmt"
	"strings"
)

type Plane struct {
	name string
}
type Car struct {
	name string
}

func (p Plane) move() string {
	return p.name
}
func (c Car) move() string {
	return c.name
}

type Speaker interface {
	move() string
}

func SpeakTwice(s Speaker) {
	fmt.Println(strings.Repeat(s.speak(), 2))
}
func main() {
	boing := Plane{"Самолет", "летит."}

	SpeakTwice(boing)
}
