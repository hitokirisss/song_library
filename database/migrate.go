package database

import (
	"database/sql"
	"log"
)

func MigrateDB(db *sql.DB) {
	createTableQuery := `
    CREATE TABLE IF NOT EXISTS songs (
        id SERIAL PRIMARY KEY,
        group_name TEXT,
        title TEXT,
        release_date TEXT,
        lyrics TEXT,
        link TEXT
    );`
	if _, err := db.Exec(createTableQuery); err != nil {
		log.Fatalf("Could not create table: %v", err)
	}
}
