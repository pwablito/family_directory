package db

func GetDiskDatabase(filename string) Database {
	return Database{
		dbType:   "sqlite3",
		filename: filename,
	}
}

func GetMemoryDatabase() Database {
	return Database{
		dbType:   "sqlite3",
		filename: ":memory:",
	}
}
