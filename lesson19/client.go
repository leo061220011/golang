package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {

	data := []byte(`{"ID": 3, "Title": "Евгений Онегин", "Author": "Пушкин"}`)
	r := bytes.NewReader(data)
	resp, err := http.Post("http://127.0.0.1:8080/books", "application/json", r)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Сервер мне ответил! Он сказал: ", resp.Body)
}
