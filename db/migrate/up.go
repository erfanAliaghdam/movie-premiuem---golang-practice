package migrate

import (
	"database/sql"
	"fmt"
)

// UP Migrate runs all necessary migrations
func UP(db *sql.DB) error {
	migrations := []string{
		// Users table
		`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE, -- unique email 
			password TEXT NOT NULL, -- hashed password
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`,

		// orders table
		`CREATE TABLE IF NOT EXISTS orders (
			id INTEGER PRIMARY KEY AUTOINCREMENT,              -- UUID for the order
			user_id INTEGER NOT NULL,         -- Foreign key referencing users table
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			paid BOOLEAN NOT NULL DEFAULT FALSE,       -- Payment status
			paid_price REAL NOT NULL DEFAULT 0.0,      -- Price after payment
			FOREIGN KEY(user_id) REFERENCES users(id)  -- Enforce user relationship
		);`,

		// licences table
		`CREATE TABLE IF NOT EXISTS licences (
    		id INTEGER PRIMARY KEY AUTOINCREMENT,
    		title TEXT,
    		finish_month INTEGER NOT NULL,
    		price REAL NOT NULL DEFAULT 0.0,
    		licence_type INTEGER NOT NULL
		);`,

		// users licence table
		`CREATE TABLE IF NOT EXISTS users_licences (
    		id INTEGER PRIMARY KEY AUTOINCREMENT,
    		user_id INTEGER NOT NULL,
    		license_id INTEGER NOT NULL,
    		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		FOREIGN KEY(user_id) REFERENCES users(id),
    		FOREIGN KEY(license_id) REFERENCES licences(id)
		);`,
	}

	for _, query := range migrations {
		if _, err := db.Exec(query); err != nil {
			return fmt.Errorf("error on executing migration: %s", err)
		}
	}

	return nil
}
