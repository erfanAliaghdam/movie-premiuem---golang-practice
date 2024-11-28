package migrate

import (
	"database/sql"
	"fmt"
)

// DOWN Migrate runs all necessary migrations
func DOWN(db *sql.DB) error {
	migrations := []string{
		// Users table
		`DROP TABLE IF EXISTS users`,
		`DROP TABLE IF EXISTS orders`,
		`DROP TABLE IF EXISTS licenses`,
		`DROP TABLE IF EXISTS users_licenses`,
	}

	for _, query := range migrations {
		if _, err := db.Exec(query); err != nil {
			return fmt.Errorf("error on executing migration: %s", err)
		}
	}

	return nil
}
