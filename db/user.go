package db

import (
	"family_directory/model"
)

func (db *Database) GetUserByUsername(username string) (*model.User, error) {
	db.Connect()
	defer db.Disconnect()
	queryString := "SELECT * FROM users WHERE username = ? LIMIT 1"
	statement, err := db.db.Prepare(queryString)

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
	db.Connect()
	defer db.Disconnect()
	insertString := "INSERT INTO user (username, name, email, password_hash, password_salt) VALUES (?, ?, ?, ?, ?)"
	statement, err := db.db.Prepare(insertString)
	if err != nil {
		return err
	}
	_, err = statement.Exec(user.Username, user.Name, user.Email, user.PasswordHash, user.PasswordSalt)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) UpdateUserName(username string, name string) error {
	db.Connect()
	defer db.Disconnect()
	updateString := "UPDATE users SET name = ? WHERE username = ?"
	statement, err := db.db.Prepare(updateString)
	if err != nil {
		return err
	}
	_, err = statement.Exec(name, username)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) UpdateUserEmail(username string, email string) error {
	db.Connect()
	defer db.Disconnect()
	updateString := "UPDATE users SET email = ? WHERE username = ?"
	statement, err := db.db.Prepare(updateString)
	if err != nil {
		return err
	}
	_, err = statement.Exec(email, username)
	if err != nil {
		return err
	}
	return nil
}
