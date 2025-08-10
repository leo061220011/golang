package main

import (
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "Index Page")
	http.ServeFile(w, r, "./static/index.html")
}
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "About Page")

}
func ContactsHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "Contacts Page")
	http.ServeFile(w, r, "./static/contacts.html")
}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/about", AboutHandler)
	http.HandleFunc("/contacts", ContactsHandler)

	fmt.Println("Server is listening...")
	http.ListenAndServe("127.0.0.1:8080", nil)
}
