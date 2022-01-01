package db

import "family_directory/model"

func (db *Database) AddChildRelationship(child model.Child, owner string) error {
	insertStatement := `
		INSERT INTO child(id, parent_id, notes, owner_username)
		VALUES (?, ?, ?, ?)
	`
	statement, err := db.db.Prepare(insertStatement)
	if err != nil {
		return err
	}
	_, err = statement.Exec(child.Id, child.ParentId, child.Notes, child.OwnerUsername)
	if err != nil {
		return err
	}
	return nil
}
