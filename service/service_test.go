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

func TestAuthServiceLogin(t *testing.T) {
	database := db.GetDiskDatabase("test.db")
	database.Create()
	defer database.DestroyIfExists()

	auth_svc := CreateAuthService(&database)

	_, err := auth_svc.RegisterUser("test", "test", "test")
	if err != nil {
		t.Error("Error registering user")
	}

	usr, err := auth_svc.LoginUser("test", "test")
	if err != nil {
		t.Error("Failed logging in after registration")
	}
	if usr == nil {
		t.Error("Returned user should be non-nil")
	}
	usr, err = auth_svc.LoginUser("test", "incorrect_password")
	if err == nil {
		t.Error("Should have failed on incorrect password")
	}
	if usr != nil {
		t.Error("Returned user should be nil")
	}

}

func TestAuthServiceValidateToken(t *testing.T) {
	database := db.GetDiskDatabase("test.db")
	database.Create()
	defer database.DestroyIfExists()

	auth_svc := CreateAuthService(&database)

	_, err := auth_svc.RegisterUser("test", "test", "test")

	if err != nil {
		t.Error("Error registering user to test token validation")
	}

	usr, err := auth_svc.LoginUser("test", "test")

	if err != nil {
		t.Error("Failed logging in after registration to test token validation")
	}

	if usr.Token == "" {
		t.Error("Token should be set after login")
	}

	if !auth_svc.ValidateToken(usr.Token, "test") {
		t.Error("Valid token verification should return true")
	}

	if auth_svc.ValidateToken("incorrect_token", "test") {
		t.Error("Invalid token verification should return false")
	}
}
