package db

import (
	"errors"
	"family_directory/model"
)

func (db *Database) GetUserByUsername(username string) (*model.User, error) {
	db.Connect()
	queryString := "SELECT * FROM users WHERE username = ? LIMIT 1"
	statement, err := db.db.Prepare(queryString)

	defer db.Disconnect()

	if err != nil {
		return nil, err
	}
	row, err := statement.Query()

	if err != nil {
		return nil, err
	}

	var result_username string
	var result_name string
	var result_email string
	var result_password_hash string
	var result_password_salt string

	row.Scan(&result_username, &result_name, &result_email, &result_password_hash, &result_password_salt)

	return &model.User{
		Username:     result_username,
		Name:         result_name,
		Email:        result_email,
		PasswordHash: result_password_hash,
		PasswordSalt: result_password_salt,
	}, nil
}

func (db *Database) AddUser(user model.User) error {
	return errors.New("not implemented")
}
