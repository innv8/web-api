package models

import (
	"database/sql"
	"log"

	"github.com/innv8/web-api/entities"
)

func CreateUser(data entities.User, db *sql.DB) (err error) {
	// 1. create a query to insert into user table
	var query = `INSERT INTO user (first_name, last_name, phone_number, date_of_birth)
		VALUES (?,?,?,?)`

	row, err := db.Exec(query, data.FirstName, data.LastName, data.PhoneNumber, data.DateOfBirth)
	if err != nil {
		log.Printf("unable to create new user because %v", err)
		return
	}

	primaryKey, _ := row.LastInsertId()
	log.Printf("created new user with id %d", primaryKey)
	return nil
}
