package db

import (
	"family_directory/model"
	"testing"
)

func TestCreateMemoryDatabase(t *testing.T) {
	database := GetMemoryDatabase()
	database.DestroyIfExists()
	err := database.Create()
	if err != nil {
		t.Errorf("Error creating database: %s", err)
	}
	// TODO make some assertions about the database
}

func TestCreateDiskDatabase(t *testing.T) {
	database := GetDiskDatabase("test.db")
	database.DestroyIfExists()
	err := database.Create()
	if err != nil {
		t.Errorf("Error creating database: %s", err)
	}
	// TODO make some assertions about the database
}

func TestCreateGetUser(t *testing.T) {
	database := GetDiskDatabase("test.db")
	database.DestroyIfExists()
	err := database.Create()
	if err != nil {
		t.Errorf("Error creating database: %s", err)
	}
	err = database.AddUser(model.User{
		Name:         "test",
		Username:     "test",
		Email:        "test",
		PasswordHash: "test",
		PasswordSalt: "test",
	})
	if err != nil {
		t.Error(err)
	}
	user, err := database.GetUserByUsername("test")
	if err != nil {
		t.Errorf("Error getting user: %s", err)
	}
	if user == nil {
		t.Errorf("User not found")
	}
	if user.Name != "test" {
		t.Errorf("User name is not correct: %s", user.Name)
	}
	if user.Email != "test" {
		t.Errorf("User email is not correct: %s", user.Email)
	}
	if user.PasswordHash != "test" {
		t.Errorf("User password hash is not correct: %s", user.PasswordHash)
	}
	if user.PasswordSalt != "test" {
		t.Errorf("User password salt is not correct: %s", user.PasswordSalt)
	}
	if user.Username != "test" {
		t.Errorf("User username is not correct: %s", user.Username)
	}
}
