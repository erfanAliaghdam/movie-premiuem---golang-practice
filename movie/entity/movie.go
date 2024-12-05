package entity

type Movie struct {
	ID          int64  `json:"id"`
	Title       string `json:"name"`
	CreatedAt   string `json:"created_at"`
	Description string `json:"description"`
}
