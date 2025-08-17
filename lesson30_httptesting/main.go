package main

import (
	"fmt"
	hd "lesson30_httptesting/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/status", hd.StatusHandler)
	http.HandleFunc("/healthy", hd.HealthHandler)

	fmt.Println("Server starting...")
	http.ListenAndServe(":8080", nil)
}
