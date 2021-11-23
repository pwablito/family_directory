package db

import (
	"testing"
)

func TestCreateMemoryDatabase(t *testing.T) {
	database := GetMemoryDatabase()
	database.Create()
}

func TestCreateDiskDatabase(t *testing.T) {
	database := GetDiskDatabase("test.db")
	database.Create()
}
