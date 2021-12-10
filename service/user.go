package service

import (
	"errors"
	"family_directory/db"
	"family_directory/model"
)

type UserService struct {
	DB *db.Database
}

func CreateUserService(database *db.Database) *UserService {
	return &UserService{
		DB: database,
	}
}

func (svc *UserService) EditUser(user *model.User) error {
	fetched_user, err := svc.DB.GetUserByUsername(user.Username)
	if err != nil {
		return err
	}
	if fetched_user == nil {
		return errors.New("user doesn't exist")
	}
	if fetched_user.Name != user.Name {
		svc.DB.UpdateUserName(user.Name, user.Name)
	}
	if fetched_user.Email != user.Email {
		svc.DB.UpdateUserEmail(user.Email, user.Name)
	}
	return nil
}
