// https://github.com/Go-cource/lesson15

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

/*
GET /books -получить все книги
GET /books/{id} - получить одну книку с идентификатором

POST /books - отправить (создать) новую книгу
*/

type Book struct {
	ID     uint
	Title  string
	Author string
}

var books = []Book{}

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, books)
	json.NewEncoder(w).Encode(books)
}

func AddBookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Add Book")
}

func main() {
	books = append(books, Book{ID: 1, Title: "Исскуство любить", Author: "Эрик Форм"})
	books = append(books, Book{ID: 2, Title: "Война и мир", Author: "Лев Толстой"})
	r := mux.NewRouter()
	r.HandleFunc("/books", BooksHandler)
	r.HandleFunc("/books/{id}", AddBookHandler)
	fmt.Println("Стартуем! Я сказал стартуем..")
	http.ListenAndServe("127.0.0.1:8080", r)

}
