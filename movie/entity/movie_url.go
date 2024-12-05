package entity

import "time"

type MovieUrl struct {
	ID        int64     `json:"id"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	Title     string    `json:"title"`
	MovieID   int64     `json:"movie_id"`
}
