package entity

import "time"

type Movie struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"image_url"`
}
