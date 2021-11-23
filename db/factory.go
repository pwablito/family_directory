package db

func GetDatabase() Database {
	return Database{
		dbType:   "sqlite3",
		filename: "db.sqlite",
	}
}

func GetTestDatabase() Database {
	return Database{
		dbType:   "sqlite3",
		filename: ":memory:",
	}
}
