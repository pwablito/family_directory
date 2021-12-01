package db

import (
	"database/sql"

	// Add database drivers here
	_ "github.com/mattn/go-sqlite3"
)

func (db *Database) Connect() error {
	db.db, _ = sql.Open(db.dbType, db.filename)
	return nil
}

func (db *Database) Disconnect() error {
	db.db.Close()
	return nil
}
