package service

import (
	"family_directory/db"
	"family_directory/model"
)

type PersonService struct {
	db *db.DB
}

func CreatePersonService() *PersonService {
	return &PersonService{db.CreateDB()}
}

func (svc *PersonService) GetPerson(id string) *db.Person {

}

func (svc *PersonService) AddPerson(person model.Person) error {
	return svc.db.AddPerson(person)
}
