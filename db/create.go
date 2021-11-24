package db

func (db *Database) Create() {
	db.CreatePersonTable()
	db.CreateChildTable()
	db.CreatePartnershipTable()
}

func (db *Database) CreatePersonTable() error {
	db.Connect()
	createSQL := `
		CREATE TABLE person (
			"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
			"name" TEXT,
			"birthdate" TEXT,
			"email" TEXT,
			"phone" TEXT
		);
	`
	statement, err := db.db.Prepare(createSQL)
	defer db.Disconnect()
	if err != nil {
		return err
	}
	statement.Exec()
	return nil
}

func (db *Database) CreateChildTable() error {
	db.Connect()
	createSQL := `
		CREATE TABLE child (
			"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
			FOREIGN KEY("childid") REFERENCES person(id),
			FOREIGN KEY("parent1id") REFERENCES person(id),
			FOREIGN KEY("parent2id") REFERENCES person(id)
		);
	`
	statement, err := db.db.Prepare(createSQL)
	defer db.Disconnect()
	if err != nil {
		return err
	}
	statement.Exec()
	return nil
}

func (db *Database) CreatePartnershipTable() error {
	db.Connect()
	createSQL := `
		CREATE TABLE partnership (
			"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
			FOREIGN KEY("person1id") REFERENCES person(id),
			FOREIGN KEY("parent2id") REFERENCES person(id),
			"start" TEXT,
			"finish" TEXT
		);
	`
	statement, err := db.db.Prepare(createSQL)
	defer db.Disconnect()
	if err != nil {
		return err
	}
	statement.Exec()
	return nil
}
