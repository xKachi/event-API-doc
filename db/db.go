package db

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

// Setting-up the database instance
var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "api.db")

	// Checking for error; if the database was not connected properly
	if err != nil {
		panic("Could not connect to database.")
	}

	/*
	Manage open database connections [Connection Pooling]

	This will enable us control the amount of connection that can be opened
	simultaneously at most, which will make sure later when the application runs we 
	don't keep on opening new connections everytime, instead we have a "pool" of ungoing
	connections that can be used whenever they are needed by different parts of the application

	maxIdleConns() — How many connections we want to keep open if the connections(max) are not
	being used at any given time.
	*/ 
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`

	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("Could not create users table")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL, 
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
		)
	`

	 _, err = DB.Exec(createEventsTable)

	 if err != nil {
		panic("Could not create events table.")
	 }
}