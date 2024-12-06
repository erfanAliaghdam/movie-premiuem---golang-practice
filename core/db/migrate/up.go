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

		// licenses table
		`CREATE TABLE IF NOT EXISTS licenses (
    		id INTEGER PRIMARY KEY AUTOINCREMENT,
    		title VARCHAR(255),
    		finish_month INTEGER NOT NULL,
    		price REAL NOT NULL DEFAULT 0.0,
    		license_type INTEGER NOT NULL
		);`,

		// users license table
		`CREATE TABLE IF NOT EXISTS users_licenses (
    		id INTEGER PRIMARY KEY AUTOINCREMENT,
    		user_id INTEGER NOT NULL,
    		license_id INTEGER NOT NULL,
    		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		FOREIGN KEY(user_id) REFERENCES users(id),
    		FOREIGN KEY(license_id) REFERENCES licenses(id)
		);`,

		// movies table
		`CREATE TABLE IF NOT EXISTS movies (
    		id INTEGER PRIMARY KEY AUTOINCREMENT,
    		title VARCHAR(255),
    		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		description Text,
    		image_url VARCHAR(500)         
    	);`,

		// movie urls table
		`CREATE TABLE IF NOT EXISTS movie_urls (
    		id INTEGER PRIMARY KEY AUTOINCREMENT,
    		title VARCHAR(255),
    		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    		url Text,
    		movie_id INTEGER NOT NULL,
    		FOREIGN KEY(movie_id) REFERENCES movies(id)
    	);`,
	}

	for _, query := range migrations {
		if _, err := db.Exec(query); err != nil {
			return fmt.Errorf("error on executing migration: %s", err)
		}
	}

	return nil
}
