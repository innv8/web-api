package utils

import (
	"database/sql"
	"log"
	"os"
)

func DBConnect() (db *sql.DB, err error) {
	// mysql connection string structure
	// username:password@tcp(host:port)/dbname
	var dbURI = os.Getenv("DB")
	db, err = sql.Open("mysql", dbURI)
	if err != nil {
		log.Printf("unable to connect to mysql because %v", err)
		return
	}

	err = db.Ping()
	if err != nil {
		log.Printf("unable to ping database because %v", err)
	}
	log.Println("connected to db")
	return db, nil
}
