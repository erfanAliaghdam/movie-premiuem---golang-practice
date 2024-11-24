package entity

type Licence struct {
	ID          string  `json:"id" db:"id"`
	Title       string  `json:"title" db:"title"`
	FinishMonth int     `json:"finish_month" db:"finish_month"`
	Price       float64 `json:"price" db:"price"`
	LicenceType int8    `json:"licence_type" db:"licence_type"`
}

const (
	RegularType  = 0
	PremiuemType = 1
)
