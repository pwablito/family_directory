package model

import (
	"errors"
	"fmt"
)

type User struct {
	Username      string
	Name          string
	Email         string
	PasswordHash  string
	PasswordSalt  string
	Token         string
	TokenLastSeen string
}

func (user *User) Serialize() string {
	return fmt.Sprintf("{\"username\":\"%s\",\"email\":\"%s\",\"name\":\"%s\"}", user.Username, user.Email, user.Name)
}

func DeserializeUser(json string) (*User, error) {
	return nil, errors.New("not implemented")
}
