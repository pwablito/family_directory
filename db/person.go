package db

import (
	"errors"
	"family_directory/model"
)

func (db *Database) AddPerson(person model.Person, owner string) error {
	insertStatement := `
		INSERT INTO person(name, birthdate, email, phone, owner, notes)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	statement, err := db.db.Prepare(insertStatement)
	if err != nil {
		return err
	}
	_, err = statement.Exec(person.Name, person.Birthdate, person.Email, person.Phone, owner, person.Notes)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) GetAllPersonsByOwner(owner string) ([]model.Person, error) {
	queryStatement := `
		SELECT * FROM person WHERE owner_username=?
	`
	statement, err := db.db.Prepare(queryStatement)
	if err != nil {
		return nil, err
	}

	row, err := statement.Query(owner)
	if err != nil {
		return nil, err
	}

	persons := make([]model.Person, 0)
	var result_id int
	var name string
	var birthdate string
	var email string
	var phone string
	var notes string
	for row.Next() {
		row.Scan(&result_id, &name, &birthdate, &email, &phone, &notes)
		persons = append(persons, model.Person{
			Id:            result_id,
			Name:          name,
			Birthdate:     birthdate,
			Email:         email,
			Phone:         phone,
			OwnerUsername: owner,
			Notes:         notes,
		})
	}
	return persons, nil
}

func (db *Database) GetPersonById(id int) (*model.Person, error) {
	queryStatement := `
		SELECT * FROM person WHERE id=?
	`
	statement, err := db.db.Prepare(queryStatement)
	if err != nil {
		return nil, err
	}

	row, err := statement.Query(id)
	if err != nil {
		return nil, err
	}

	var result_id int
	var name string
	var birthdate string
	var email string
	var phone string
	var owner string
	var notes string

	row.Scan(&result_id, &name, &birthdate, &email, &phone, &owner, &notes)

	return &model.Person{
		Id:            result_id,
		Name:          name,
		Birthdate:     birthdate,
		Email:         email,
		Phone:         phone,
		OwnerUsername: owner,
		Notes:         notes,
	}, nil
}

func (db *Database) RemovePerson(id int) error {
	deleteStatement := `
		DELETE FROM user WHERE id=?
	`
	statement, err := db.db.Prepare(deleteStatement)
	if err != nil {
		return err
	}
	_, err = statement.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) UpdatePerson(person model.Person) error {
	return errors.New("not implemented")
}
