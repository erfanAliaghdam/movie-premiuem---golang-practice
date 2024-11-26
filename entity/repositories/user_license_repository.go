package repositories

import (
	"database/sql"
	"fmt"
)

type UserLicenseRepository interface {
	CreateUserLicense(licenseID int64, userID int64) (bool, error)
	CheckIfUserHasActiveLicense(userID int64) (bool, error)
}

type userLicenseRepository struct {
	db *sql.DB
}

func NewUserLicenseRepository(db *sql.DB) UserLicenseRepository {
	return &userLicenseRepository{db: db}
}

func (r *userLicenseRepository) CreateUserLicense(licenseID int64, userID int64) (bool, error) {
	query := `INSERT INTO users_licenses(user_id, license_id) VALUES (?, ?)`

	_, err := r.db.Exec(query, userID, licenseID)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *userLicenseRepository) CheckIfUserHasActiveLicense(userID int64) (bool, error) {
	// SQL query to check for active licenses
	query := `
        SELECT COUNT(*)
        FROM users_licenses ul
        JOIN licenses l ON ul.license_id = l.id
        WHERE ul.user_id = ?
          AND DATETIME(ul.created_at, '+' || l.finish_month || ' months') > DATETIME('now');
    `

	// Placeholder for the count result
	var count int

	// Execute the query
	err := r.db.QueryRow(query, userID).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error checking active license: %w", err)
	}

	// Return true if the count is greater than 0, meaning the user has an active license
	return count > 0, nil
}
