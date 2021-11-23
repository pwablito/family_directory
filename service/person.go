package service

import (
	"family_directory/db"
	"family_directory/model"
)

type PersonService struct {
	db db.Database
}

func CreatePersonService() *PersonService {
	return &PersonService{
		db: db.GetDatabase(),
	}
}

func (svc *PersonService) GetPersonById(id int) (*model.Person, error) {
	person, err := svc.db.GetPersonById(id)
	if err != nil {
		return nil, err
	}
	return person, nil
}

func (svc *PersonService) AddPerson(person model.Person) error {
	return svc.db.AddPerson(person)
}
