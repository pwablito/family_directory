package service

import (
	"errors"
	"family_directory/db"
	"family_directory/model"
	"family_directory/util"
)

type AuthService struct {
	DB *db.Database
}

func CreateAuthService(db *db.Database) *AuthService {
	return &AuthService{DB: db}
}

func (svc *AuthService) RegisterUser(username string, password string, email string) (*model.User, error) {
	user, err := svc.DB.GetUserByUsername(username)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			return nil, err
		}
	}
	if user != nil {
		return nil, errors.New("username already taken")
	}
	salt := util.RandomString(10)
	hash := HashPassword(password, salt)
	user = &model.User{
		Username:     username,
		Email:        email,
		PasswordHash: hash,
		PasswordSalt: salt,
	}
	err = svc.DB.AddUser(*user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (svc *AuthService) LoginUser(username string, password string) (*model.User, error) {
	user, err := svc.DB.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not found")
	}
	if !VerifyPassword(password, user.PasswordSalt, user.PasswordHash) {
		return nil, errors.New("invalid password")
	}
	return user, nil // TODO add some sort of session token
}

func HashPassword(password string, salt string) string {
	return util.HashString(salt + password)
}

func VerifyPassword(password string, salt string, hash string) bool {
	return HashPassword(salt, password) == hash
}

func (svc *AuthService) VerifyToken(token string, username string) bool {
	user, err := svc.DB.GetUserByUsername(username)
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
	svc.DB.SetTokenForUser(username, new_token, util.GetCurrentTime())
	return new_token
}
