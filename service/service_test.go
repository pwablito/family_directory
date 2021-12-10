package service

import (
	"family_directory/db"
	"testing"
)

func TestAuthServiceRegister(t *testing.T) {
	database := db.GetDiskDatabase("test.db")
	database.Create()
	defer database.DestroyIfExists()

	auth_svc := CreateAuthService(&database)

	usr, err := auth_svc.RegisterUser("test", "test", "test")
	if err != nil {
		t.Error("Error registering user")
	}
	if usr.Username != "test" {
		t.Error("Username not set correctly")
	}
	if usr.Email != "test" {
		t.Error("Email not set correctly")
	}

	_, err = auth_svc.RegisterUser("test", "test", "test")
	if err == nil {
		t.Error("User should not be able to register twice")
	}
}
