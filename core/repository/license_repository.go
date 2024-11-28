package repository

import (
	"database/sql"
	"movie_premiuem/core/entity"
)

type LicenseRepository interface {
	CreateLicense(license entity.License) (entity.License, error)
	GetAllLicenses() ([]entity.License, error)
}

type licenseRepository struct {
	db *sql.DB
}

func NewLicenseRepository(db *sql.DB) LicenseRepository {
	return &licenseRepository{db: db}
}

func (r *licenseRepository) CreateLicense(license entity.License) (entity.License, error) {
	query := `INSERT INTO licenses (title, finish_month, license_type, price) VALUES (?, ?, ?, ?)`

	result, err := r.db.Exec(query, license.Title, license.FinishMonth, license.LicenseType, license.Price)
	if err != nil {
		return entity.License{}, err
	}

	id, _ := result.LastInsertId()
	license.ID = id
	return license, nil
}

func (r *licenseRepository) GetAllLicenses() ([]entity.License, error) {
	var licenses []entity.License
	query := `SELECT id, title, finish_month, price, license_type FROM licenses;`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var license entity.License
		rowErr := rows.Scan(&license.ID, &license.Title, &license.FinishMonth, &license.Price, &license.LicenseType)
		if rowErr != nil {
			return nil, rowErr
		}
		licenses = append(licenses, license)
	}

	if rowsErr := rows.Err(); rowsErr != nil {
		return nil, rowsErr
	}

	return licenses, nil
}
