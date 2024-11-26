package repositories

import (
	"database/sql"
	"movie_premiuem/entity"
)

type LicenseRepository interface {
	CreateLicense(license entity.License) (entity.License, error)
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
