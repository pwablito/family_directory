package main

import (
	handler "family_directory/handler"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server...")
	http.HandleFunc("/add_person", handler.AddPerson)
	http.ListenAndServe(":8080", nil)
}
