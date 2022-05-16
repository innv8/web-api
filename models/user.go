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

func ListUsers(firstName string, db *sql.DB) (users []entities.User, err error) {
	users = []entities.User{}
	var user entities.User
	var _updated sql.NullString

	var query = `SELECT id, first_name, last_name, phone_number, date_of_birth, created, updated 
		FROM user`
	if firstName != "" {
		query += " WHERE first_name LIKE '%" + firstName + "%'"
	}
	log.Println(query)

	rows, err := db.Query(query)
	if err != nil {
		log.Printf("unable to fetch users because %v", err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&user.ID, &user.FirstName, &user.LastName, &user.PhoneNumber, &user.DateOfBirth,
			&user.Created, &_updated)
		if err != nil {
			log.Printf("unable to scan user because %v", err)
			return nil, err
		}
		user.Updated = _updated.String
		users = append(users, user)
	}
	log.Printf("got %d users from the db", len(users))
	return users, nil
}
