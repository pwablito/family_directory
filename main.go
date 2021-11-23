package main

import (
	"family_directory/db"
	"family_directory/handler"
	"fmt"
	"net/http"
)

func main() {
	database := db.GetTestDatabase()
	database.Create()
	fmt.Println("Starting server...")
	http.HandleFunc("/add_person", handler.AddPerson)
	http.ListenAndServe(":8080", nil)
}
