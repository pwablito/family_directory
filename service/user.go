package service

import (
	"errors"
	"family_directory/db"
	"family_directory/model"
)

type UserService struct {
	db db.Database
}

func CreateUserService() *UserService {
	return &UserService{}
}

func (svc *UserService) EditUser(user *model.User) error {
	fetched_user, err := svc.db.GetUserByUsername(user.Username)
	if err != nil {
		return err
	}
	if fetched_user == nil {
		return errors.New("user doesn't exist")
	}
	if fetched_user.Name != user.Name {
		svc.db.UpdateUserName(user.Name, user.Name)
	}
	if fetched_user.Email != user.Email {
		svc.db.UpdateUserEmail(user.Email, user.Name)
	}
	return nil
}
