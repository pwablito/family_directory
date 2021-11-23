package db

import (
	"database/sql"
	"errors"
)

func (db *Database) connect() error {
	db.db, _ = sql.Open("sqlite3", db.filename)
	return nil
}

func (db *Database) disconnect() error {
	db.db.Close()
	return nil
}

func (db *Database) runQuery(q string, args ...interface{}) ([]string, error) {
	err := db.connect()
	if err != nil {
		return nil, errors.New("Connect error")
	}

	statement, err := db.db.Prepare(q)

	if err != nil {
		return nil, errors.New("Prepare SQL error")
	}

	row, err := statement.Exec(args...)

	if err != nil {
		return nil, errors.New("Exec SQL error")
	}

	defer db.disconnect()

	var rows []string

	for row.Next() {
		rows = append(rows, row)
	}
	return rows, nil
}
