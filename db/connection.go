package db

import (
	"database/sql"
)

func (db *Database) connect() error {
	db.db, _ = sql.Open(db.dbType, db.filename)
	return nil
}

func (db *Database) disconnect() error {
	db.db.Close()
	return nil
}
