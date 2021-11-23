package db

import (
	"family_directory/model"
)

func (db *Database) AddPerson(person model.Person) error {
	inserStatement := `
		INSERT INTO person(name, birthdate, email, phone)
		VALUES (?, ?, ?, ?)
	`
	statement, err := db.db.Prepare(inserStatement)
	// This is good to avoid SQL injections
	if err != nil {
		return err
	}
	_, err = statement.Exec(person.Name, person.Birthdate, person.Email, person.Phone)
	if err != nil {
		return err
	}
	return nil
}
