package model

type User struct {
	Username     string
	Email        string
	PasswordHash string
	PasswordSalt string
}
