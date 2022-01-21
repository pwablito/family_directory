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
	err = db.CreatePartnershipTables()
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
			FOREIGN KEY("owner") REFERENCES user("username")
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
			FOREIGN KEY("child_id") REFERENCES person("id"),
			FOREIGN KEY("parent_id") REFERENCES person("id"),
			FOREIGN KEY("owner") REFERENCES user("username")
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

func (db *Database) CreatePartnershipTables() error {
	db.Connect()
	defer db.Disconnect()
	createSQL := `
		CREATE TABLE partnership (
			"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
			"owner" TEXT NOT NULL,
			"start" TEXT,
			"finish" TEXT,
			"notes" TEXT,
			FOREIGN KEY("owner") REFERENCES user("username")
		);
	`
	statement, err := db.db.Prepare(createSQL)
	if err != nil {
		return err
	}
	statement.Exec()
	createSQL = `
		CREATE TABLE partnership_member (
			"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
			"partnership_id" INTEGER NOT NULL,
			"person_id" INTEGER NOT NULL,
			"start" TEXT,
			"finish" TEXT,
			"notes" TEXT,
			FOREIGN KEY("person_id") REFERENCES person("id")
		);
	`
	statement, err = db.db.Prepare(createSQL)
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
