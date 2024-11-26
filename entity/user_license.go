package entity

import "time"

type UserLicense struct {
	ID        int64     `json:"id" db:"id"`
	UserID    int64     `json:"user_id" db:"user_id"`
	LicenseID int64     `json:"license_id" db:"license_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
