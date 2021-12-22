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

func (svc *PersonService) RemovePersonById(id int, token string) error {
	owner, err := svc.db.GetUserByToken(token)
	if err != nil {
		return err
	}
	if owner == nil {
		return errors.New("invalid token")
	}
	auth_svc := CreateAuthService(&svc.db)
	if auth_svc.ValidateToken(token, owner.Username) {
		existing_entry, err := svc.db.GetPersonById(id)
		if err != nil {
			return err
		}
		if existing_entry == nil {
			return errors.New("person not found")
		}
		err = svc.db.RemovePerson(id)
		if err != nil {
			return err
		}
	}
	return errors.New("invalid token")
}

func (svc *PersonService) UpdatePerson(person model.Person, token string) error {
	// Username is used to find the user to update, so it cannot be updated
	owner, err := svc.db.GetUserByToken(token)
	if err != nil {
		return err
	}
	if owner == nil {
		return errors.New("invalid token")
	}
	auth_svc := CreateAuthService(&svc.db)
	if auth_svc.ValidateToken(token, owner.Username) {
		existing_entry, err := svc.db.GetPersonById(person.Id)
		if err != nil {
			return err
		}
		if existing_entry == nil {
			return errors.New("person not found")
		}
		err = svc.db.UpdatePerson(person)
		if err != nil {
			return err
		}
	}
	return errors.New("invalid token")
}
