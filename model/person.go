package model

import (
	"time"
)

type Person struct {
	Id        int
	Name      string
	Birthdate time.Time
	Email     string
	Phone     string
}
