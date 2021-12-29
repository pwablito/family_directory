package db

func (db *Database) Create() error {
	err := db.CreateUserTable()
	if err != nil {
		return err
	}
	err = db.CreatePersonTable()
	if err != nil {
		return err
	}
	err = db.CreateChildTable()
	if err != nil {
		return err
	}
	err = db.CreatePartnershipTable()
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) DestroyIfExists() {
	db.Connect()
	db.db.Exec("DROP TABLE IF EXISTS person")
	db.db.Exec("DROP TABLE IF EXISTS child")
	db.db.Exec("DROP TABLE IF EXISTS partnership")
	db.db.Exec("DROP TABLE IF EXISTS user")
	db.Disconnect()
}

func (db *Database) CreatePersonTable() error {
	db.Connect()
	createSQL := `
		CREATE TABLE person (
			"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
			"name" TEXT,
			"birthdate" TEXT,
			"email" TEXT,
			"phone" TEXT,
			"owner" TEXT NOT NULL,
			"notes" TEXT,
			FOREIGN KEY("owner") REFERENCES user(username)
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
			"child_id" INTEGER NOT NULL,
			"parent_id" INTEGER NOT NULL,
			"owner" TEXT NOT NULL,
			"notes" TEXT,
			FOREIGN KEY("child_id") REFERENCES person(id),
			FOREIGN KEY("parent_id") REFERENCES person(id),
			FOREIGN KEY("owner") REFERENCES user(username)
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
			"person1_id" INTEGER NOT NULL,
			"person2_id" INTEGER NOT NULL,
			"owner" TEXT NOT NULL,
			"start" TEXT,
			"finish" TEXT,
			"notes" TEXT,
			FOREIGN KEY("person1_id") REFERENCES person(id),
			FOREIGN KEY("person2_id") REFERENCES person(id),
			FOREIGN KEY("owner") REFERENCES user(username)
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

func (db *Database) CreateUserTable() error {
	db.Connect()
	createSQL := `
		CREATE TABLE user (
			"username" TEXT NOT NULL PRIMARY KEY,
			"name" TEXT,
			"email" TEXT,
			"password_hash" TEXT,
			"password_salt" TEXT,
			"token" TEXT,
			"token_created" TEXT,
			"notes" TEXT
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
