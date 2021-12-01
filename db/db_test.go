package db

import (
	"testing"
)

func TestCreateMemoryDatabase(t *testing.T) {
	database := GetMemoryDatabase()
	database.Create()
	// TODO make some assertions about the database
}

func TestCreateDiskDatabase(t *testing.T) {
	database := GetDiskDatabase("test.db")
	database.Create()
	// TODO make some assertions about the database
}
