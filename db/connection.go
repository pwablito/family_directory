package db

import (
	"database/sql"
)

func (db *Database) Connect() error {
	db.db, _ = sql.Open(db.dbType, db.filename)
	return nil
}

func (db *Database) Disconnect() error {
	db.db.Close()
	return nil
}
