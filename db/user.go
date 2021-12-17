package db

import (
	"family_directory/model"
)

func (db *Database) GetUserByUsername(username string) (*model.User, error) {
	db.Connect()
	defer db.Disconnect()
	queryString := "SELECT username, name, email, password_hash, password_salt, token, token_created FROM user WHERE username = $1 LIMIT 1"

	row := db.db.QueryRow(queryString, username)

	var result_username string
	var result_name string
	var result_email string
	var result_password_hash string
	var result_password_salt string
	var result_token string
	var result_token_created string

	err := row.Scan(
		&result_username, &result_name, &result_email, &result_password_hash,
		&result_password_salt, &result_token, &result_token_created,
	)

	if err != nil {
		return nil, err
	}

	return &model.User{
		Username:     result_username,
		Name:         result_name,
		Email:        result_email,
		PasswordHash: result_password_hash,
		PasswordSalt: result_password_salt,
		Token:        result_token,
		TokenCreated: result_token_created,
	}, nil
}

func (db *Database) AddUser(user model.User) error {
	db.Connect()
	defer db.Disconnect()
	insertString := "INSERT INTO user (username, name, email, password_hash, password_salt, token, token_created) VALUES (?, ?, ?, ?, ?, ?, ?)"
	statement, err := db.db.Prepare(insertString)
	if err != nil {
		return err
	}
	_, err = statement.Exec(user.Username, user.Name, user.Email, user.PasswordHash, user.PasswordSalt, user.Token, user.TokenCreated)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) UpdateUserName(username string, name string) error {
	db.Connect()
	defer db.Disconnect()
	updateString := "UPDATE user SET name = ? WHERE username = ?"
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
	updateString := "UPDATE user SET email = ? WHERE username = ?"
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

func (db *Database) SetToken(username string, token string, timestamp string) error {
	db.Connect()
	defer db.Disconnect()
	updateString := "UPDATE user SET token = ?, token_created = ? WHERE username = ?"
	statement, err := db.db.Prepare(updateString)
	if err != nil {
		return err
	}
	_, err = statement.Exec(token, timestamp, username)
	if err != nil {
		return err
	}
	return nil
}
