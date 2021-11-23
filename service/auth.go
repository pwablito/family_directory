package service

import "family_directory/db"

type AuthService struct {
	db db.Database
}

func (svc *AuthService) RegisterUser(username string, password string, email string)
