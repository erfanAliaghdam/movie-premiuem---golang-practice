package migrate

import (
	"database/sql"
	"fmt"
)

// UP Migrate runs all necessary migrations
func UP(db *sql.DB) error {
	migrations := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`,
	}

	for _, query := range migrations {
		if _, err := db.Exec(query); err != nil {
			return fmt.Errorf("error on executing migration: %s", err)
		}
	}

	return nil
}
