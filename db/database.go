package db

import (
	"database/sql"
)

type Database struct {
	dbType   string
	filename string
	db       *sql.DB
}
