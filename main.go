package main

import (
	"family_directory/handler"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Registering handlers...")
	// Auth handlers
	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/register", handler.Register)
	// Singleton directory handler
	http.HandleFunc("/get_directory", handler.GetDirectory)
	// Person handlers
	http.HandleFunc("/get_person", handler.GetPerson)
	http.HandleFunc("/add_person", handler.AddPerson)
	http.HandleFunc("/delete_person", handler.DeletePerson)
	http.HandleFunc("/edit_person", handler.EditPerson)
	// Relationship handlers
	http.HandleFunc("/add_child_relationship", handler.AddChildRelationship)
	http.HandleFunc("/add_partner_relationship", handler.AddPartnerRelationship)
	http.HandleFunc("/delete_child_relationship", handler.DeleteChildRelationship)
	http.HandleFunc("/delete_partner_relationship", handler.DeletePartnerRelationship)
	http.HandleFunc("/edit_child_relationship", handler.EditChildRelationship)
	http.HandleFunc("/edit_partner_relationship", handler.EditPartnerRelationship)

	// Start server
	fmt.Println("Starting server...")
	http.ListenAndServe(":8080", nil)
}
