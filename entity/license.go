package entity

type License struct {
	ID          int64   `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	FinishMonth int     `json:"finish_month" db:"finish_month"`
	Price       float64 `json:"price" db:"price"`
	LicenseType int8    `json:"license_type" db:"license_type"`
}

const (
	RegularLicenseType  = 0
	PremiuemLicenseType = 1
)
