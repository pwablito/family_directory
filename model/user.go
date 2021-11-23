package model

import (
	"fmt"
)

type User struct {
	Username     string
	Name         string
	Email        string
	PasswordHash string
	PasswordSalt string
}

func (user *User) serialize() string {
	return fmt.Sprintf("{\"username\":\"%s\",\"email\":\"%s\",\"name\":\"%s\"}", user.Username, user.Email, user.Name)
}
