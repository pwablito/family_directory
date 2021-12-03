package service

import (
	"errors"
	"family_directory/db"
	"family_directory/model"
	"family_directory/util"
)

type AuthService struct {
	db db.Database
}

func (svc *AuthService) RegisterUser(username string, password string, email string) (*model.User, error) {
	user, err := svc.db.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return nil, errors.New("username already taken")
	}
	salt := util.RandomString(10)
	hash := hashPassword(password, salt)
	user = &model.User{
		Username:     username,
		Email:        email,
		PasswordHash: hash,
		PasswordSalt: salt,
	}
	err = svc.db.AddUser(*user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (svc *AuthService) LoginUser(username string, password string) (*model.User, error) {
	user, err := svc.db.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	if !verifyPassword(password, user.PasswordSalt, user.PasswordHash) {
		return nil, errors.New("invalid password")
	}
	return user, nil // TODO add some sort of session token
}

func hashPassword(password string, salt string) string {
	return util.HashString(salt + password)
}

func verifyPassword(password string, salt string, hash string) bool {
	return hashPassword(salt, password) == hash
}

func (svc *AuthService) VerifyToken(token string, username string) bool {
	user, err := svc.db.GetUserByUsername(username)
	if err != nil {
		return false
	}
	if user == nil {
		return false
	}
	if user.Token != token {
		return false
	}
	// TODO check against timestamp from database
	return true
}

func (svc *AuthService) NewTokenForUser(username string) string {
	new_token := util.RandomString(10)
	svc.db.SetTokenForUser(username, new_token, util.GetCurrentTime())
	return new_token
}
