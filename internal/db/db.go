package db

import (
	"database/sql"

	_"github.com/mattn/go-sqlite3"
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
	createToursTable := `
		CREATE TABLE IF NOT EXISTS tours (
			id TEXT PRIMARY KEY,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			dateTime DATETIME NOT NULL,
			userId VARCHAR
		);
	`
	_, err := DB.Exec(createToursTable)
	if err != nil {
		panic("Could not create tables.")
	}
}