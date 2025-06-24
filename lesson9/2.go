package main

import "fmt"

func Say(animal string) (v string) {
	switch animal {
	default:
		v = "heh"
	case "dog":
		v = "gav"
	case "cat":
		v = "myau"
	case "cow":
		v = "mu"
	}
	return
}
func Print(who string, how func(string) string) {
	fmt.Println(how(who))
}
func main() {
	var voice func(string) string
	voice = Say
	Print("cat", voice)
	fmt.Println(Say("cat"))
}
