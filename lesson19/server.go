// https://github.com/Go-cource/lesson15

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

/*
GET /books -получить все книги
GET /books/{id} - получить одну книку с идентификатором

POST /books - отправить (создать) новую книгу
*/

type Book struct {
	ID     int
	Title  string
	Author string
}

var books = []Book{}

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, books)
	//json.NewEncoder(w).Encode(books)
	if r.Method == http.MethodGet {
		json.NewEncoder(w).Encode(books)

	}
	if r.Method == http.MethodPost {
		fmt.Fprint(w, "POST AA!!!")
	}
}

func AddBookHandler(w http.ResponseWriter, r *http.Request) {
	//	fmt.Fprint(w, "Add Book")
	params := mux.Vars(r)
	//fmt.Fprint(w, params["id"])
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println(err)
	}
	for _, value := range books {
		if value.ID == id {
			json.NewEncoder(w).Encode(value)
			return

		}
	}
	json.NewEncoder(w).Encode(Book{})
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

// //*

// package main
// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"strconv"
// 	"github.com/gorilla/mux"
// )
// /*
// GET /books - получить все книги
// GET /books/{id} - получить одну книгу с идентификатором
// POST /books - отправить (создать) новую книгу
// */
// type Book struct {
// 	ID     int
// 	Title  string
// 	Author string
// }
// var books = []Book{}
// func BooksHandler(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Fprint(w, books)
// 	json.NewEncoder(w).Encode(books)
// }
// func AddBookHandler(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r) //достать все параметры из url
// 	// fmt.Fprint(w, params["id"])
// 	id, err := strconv.Atoi(params["id"]) //перевожу параметр id в число (тк в структуре у меня ID - число)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	for _, value := range books { //иду по слайсу books и сравниваю параметр id с идентификатором книги
// 		if value.ID == id {
// 			json.NewEncoder(w).Encode(value)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(Book{}) //если не найден - вывожу пустую структуру книги
// }
// func main() {
// 	books = append(books, Book{ID: 1, Title: "Искусство любить", Author: "Эрих Фромм"})
// 	books = append(books, Book{ID: 2, Title: "Война и Мир", Author: "Толстой"})
// 	r := mux.NewRouter()
// 	r.HandleFunc("/books", BooksHandler)
// 	r.HandleFunc("/books/{id}", AddBookHandler)
// 	fmt.Println("СТАРТУЕМ! Я сказал стартуем...")
// 	http.ListenAndServe("127.0.0.1:8080", r)
// }
// *//
