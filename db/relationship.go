package db

import (
	"errors"
	"family_directory/model"
)

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

func (db *Database) AddPartnershipRelationship(partnership model.Partnership, owner string) error {
	return errors.New("not implemented")
}

func (db *Database) RemoveChildRelationship(id int) error {
	deleteStatement := `
		DELETE FROM child WHERE id=?
	`
	statement, err := db.db.Prepare(deleteStatement)
	if err != nil {
		return err
	}
	_, err = statement.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (db *Database) RemovePartnershipRelationship(id int) error {
	return errors.New("not implemented")
}

func (db *Database) RemovePartnerFromRelationship(partnership_id int, person_id int) error {
	deleteStatement := `
		DELETE FROM partnership_member WHERE person_id=? AND partnership_id=?
	`
	statement, err := db.db.Prepare(deleteStatement)
	if err != nil {
		return err
	}
	_, err = statement.Exec(person_id, partnership_id)
	if err != nil {
		return err
	}
	return nil
}

/*
* Marks a partner's membership in a relationship as ended at a provided timestamp
 */
func (db *Database) EndPartnerRelationshipMembership(partnership_id int, person_id int, timestamp string) error {
	updateStatement := `
		UPDATE partnership_member SET finish=? WHERE person_id=? AND partnership_id=?
	`
	statement, err := db.db.Prepare(updateStatement)
	if err != nil {
		return err
	}
	_, err = statement.Exec(timestamp, person_id, partnership_id)
	if err != nil {
		return err
	}
	return nil
}
