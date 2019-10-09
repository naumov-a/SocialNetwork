package main

import (
	"database/sql"
	"log"
	"time"
)

type DatabaseConnection struct {
	connection *sql.DB
}

func (dbConn DatabaseConnection) init() {
	username := ""
	password := ""
	address := ""
	dbName := ""
	dataSource := username + ":" + password + "@(" + address + ")/" + dbName

	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	dbConn.connection = db
}

func (dbConn DatabaseConnection) insertUser(User) {

}

func (dbConn DatabaseConnection) getUser(id int) (User) {
	var (
		firstName string
		lastName string
		birthDate time.Time
		gender Gender
		city string
		interests []string
		password string
	)

	query := "SELECT username, password, created_at FROM users WHERE id = ?"
	if err := dbConn.connection.QueryRow(query, id).Scan(&firstName, &lastName, &birthDate, &gender, &city, &interests, &password); err != nil {
		log.Fatal(err)
	}

	return User {
		id,
		firstName,
		lastName,
		birthDate,
		gender,
		city,
		interests,
		password,
	}
}
