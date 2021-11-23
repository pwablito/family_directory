package model

import (
	"time"
)

type Person struct {
	Id       int
	Name     string
	Birthday time.Time
	Email    string
	Phone    string
}
