package db

import (
	"family_directory/model"
)

func (db *Database) AddPerson(person model.Person, owner string) error {
	insertStatement := `
		INSERT INTO person(name, birthdate, email, phone, owner)
		VALUES (?, ?, ?, ?, ?)
	`
	statement, err := db.db.Prepare(insertStatement)
	if err != nil {
		return err
	}
	_, err = statement.Exec(person.Name, person.Birthdate, person.Email, person.Phone, owner)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) GetPersonById(id int) (*model.Person, error) {
	queryStatement := `
		SELECT * FROM person WHERE id=?
	`
	statement, err := db.db.Prepare(queryStatement)
	if err != nil {
		return nil, err
	}

	row, err := statement.Query()
	if err != nil {
		return nil, err
	}

	var result_id int
	var name string
	var birthdate string
	var email string
	var phone string
	row.Scan(&result_id, &name, &birthdate, &email, &phone)

	return &model.Person{
		Id:        result_id,
		Name:      name,
		Birthdate: birthdate,
		Email:     email,
		Phone:     phone,
	}, nil
}
