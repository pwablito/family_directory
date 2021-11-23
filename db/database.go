package db

type Database struct {
	filename string
	db       sql.DB
}
