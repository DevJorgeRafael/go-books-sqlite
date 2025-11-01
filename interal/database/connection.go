package database

import (
	"database/sql"
	"fmt"
)

type Config struct {
	DatabasePath string
}

func New(cfg Config) (*sql.DB, error) {
	db, err := sql.Open("sqlite", cfg.DatabasePath)
	if err != nil {
		return nil, fmt.Errorf("error abriendo base de datos: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error conectando a la base de datos: %w", err)
	}

	return db, nil
}

func Close(db *sql.DB) error {
	if err := db.Close(); err != nil {
		return fmt.Errorf("error cerrando base de datos: %w", err)
	}
	return nil
}