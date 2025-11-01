package database

import (
	"database/sql"
	"fmt"
)

func RunMigrations(db *sql.DB) error {
	migrations := []string{
		createBooksTable,
		createAuthorsTable,
		createAuthorBooksTable,
	}

	for i, migration := range migrations {
		if _, err := db.Exec(migration); err != nil {
			return fmt.Errorf("error ejecutando migración %d: %w", i+1, err)
		}
	}
	return nil
}

// Cada constante es una migracion SQL
const createBooksTable = `
	CREATE TABLE IF NOT EXISTS books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		publication_year INTEGER,
		isbn TEXT UNIQUE,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)
`

const createAuthorsTable = `
	CREATE TABLE IF NOT EXISTS authors (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		biography TEXT,
		country TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)
`

const createAuthorBooksTable = `
	CREATE TABLE IF NOT EXISTS author_books (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		author_id INTEGER NOT NULL,
		book_id INTEGER NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

		FOREIGN KEY (author_id) REFERENCES authors(id) ON DELETE CASCADE,
		FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE,

		UNIQUE (author_id, book_id)
	)
		
`