package main

import "fmt"

func Divide(a, b float64) (float64, error) {
	if b != 0 {
		return a / b, nil
	}
	return 0, fmt.Errorf("Ошибка! Деление на 0")
}
func main() {
	result, err := Divide(10, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

/* команда для компиляции в Windows: GOOS=windows GOARCH=amd64 go build -o main_for_Windows.exe

команда для компиляции в Lunix:  $env:GOARCH="amd64"; go build -o 1_for_linux 1.go
*/
