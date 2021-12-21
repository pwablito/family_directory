package service

import (
	"errors"
	"family_directory/db"
	"family_directory/model"
)

type PersonService struct {
	db db.Database
}

func CreatePersonService() *PersonService {
	return &PersonService{}
}

func (svc *PersonService) GetPersonById(id int) (*model.Person, error) {
	person, err := svc.db.GetPersonById(id)
	if err != nil {
		return nil, err
	}
	return person, nil
}

func (svc *PersonService) AddPerson(person model.Person, token string) error {
	owner, err := svc.db.GetUserByToken(token)
	if err != nil {
		return err
	}
	if owner == nil {
		return errors.New("invalid token")
	}
	auth_svc := CreateAuthService(&svc.db)
	if auth_svc.ValidateToken(token, owner.Username) {
		svc.db.AddPerson(person, owner.Username)
	}
	return errors.New("invalid token")
}
