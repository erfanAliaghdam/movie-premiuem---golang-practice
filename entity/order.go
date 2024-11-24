package entity

import "time"

type Order struct {
	ID        int64     `json:"id" db:"id"`                 // UUID for the order
	UserID    int64     `json:"user_id" db:"user_id"`       // Name or identifier of the customer
	CreatedAt time.Time `json:"created_at" db:"created_at"` // Timestamp when the order was created
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"` // Timestamp when the order was last updated
	Paid      bool      `json:"paid" db:"paid"`             // true if payment was success
	PaidPrice float64   `json:"paid_price" db:"paid_price"` // price will be saved after payment becomes true
}
