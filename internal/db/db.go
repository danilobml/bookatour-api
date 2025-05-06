package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id TEXT PRIMARY KEY,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL,
			role TEXT
		);
	`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Could not create user table.")
	}

	createToursTable := `
		CREATE TABLE IF NOT EXISTS tours (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			dateTime DATETIME NOT NULL,
			userId TEXT NOT NULL,
			FOREIGN KEY(userId) REFERENCES users(id)
		);
	`
	_, err = DB.Exec(createToursTable)
	if err != nil {
		panic("Could not create tours table.")
	}

	createBookingsTable := `
		CREATE TABLE IF NOT EXISTS bookings (
			id TEXT PRIMARY KEY,
			tourId TEXT NOT NULL,
			userId TEXT NOT NULL,
			FOREIGN KEY(tourId) REFERENCES tours(id),
			FOREIGN KEY(userId) REFERENCES users(id)
		)
	`

	_, err = DB.Exec(createBookingsTable)
	if err != nil {
		panic("Could not create bookings table.")
	}
}
