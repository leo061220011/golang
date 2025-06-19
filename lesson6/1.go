/*package main

func main() {

	/*for i := 0; i < 255; i++ {
		fmt.Printf("%d - %c \n", i, string(byte(i)))
	}*/
/*
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33
	fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang)
*/
/*
	myRune := 'A'            // = 65
	fmt.Printf("%c", myRune) // 65 -> A
*/

/*
	myRune := 128515
	fmt.Printf("%c", myRune)
*/
/*
	message := "Hello!!"
	fmt.Println(message[0])
	fmt.Printf("%c", message[0])
	message[0] = 'T' // error! строки неизменяемы
*/
/*
	myString := "Hello!"

	for i := 0; i < 6; i++ {
		fmt.Printf("%c\n", myString[i])

	}
*/

/*// кодирование
c := 'a'
c = c + 3 // c += 3
if c > 'z' {
	c = c - 26
}
fmt.Printf("%c", c)
*/

/*
	c := 'a'
	c = c + 3 // c += 3
	if c > 'z' {
		c = c - 26
	}
	fmt.Printf("%c", c)
*/
/*
	c := 'g' // как работает unicod таблица шифр Цезаря
	c = c - 'b' + 'B'
	fmt.Printf("%c", c)

*/
/*message := "uv vargeangvbany fcnpr fgngvba"
fmt.Println(len(message))*/
//myLen := len(message)
/*
		message := "Привет"
		fmt.Println(utf8.RuneCountInString(message))
}
*/
/*
package main

import (
	"fmt"
)

func main() {

	/*message := "uv vagreangvbany fcnpr fgngvba"
	messageLength := len(message)
	for i := 0; i < messageLength; i++ {
		c := message[i]
		if c >= 'a' && c <= 'z' {
			c += 13
			if c > 'z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
	}
*/

/*question := "¿Cómo estás?"
fmt.Println(len(question), "bytes")
fmt.Println(utf8.RuneCountInString(question), "runes")
c, size := utf8.DecodeRuneInString(question)
fmt.Printf("First rune: %c %v bytes", c, size)
*/

/*question := "¿Cómo estás?"
	for i, c := range question { // цикл с использованием rangeб но надо использовать  2 переменные. но лишнюю переменную можно заменить нижним подчеркиванием i -> _
		fmt.Printf("%v %c\n", i, c)
	}

}
*/