package db

func GetDatabase() Database {
	return Database{
		filename: "db.sqlite",
	}
}
