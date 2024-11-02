package models

type Image struct {
	ID         int    `json:"id"`
	MovieID    int    `json:"movie_id"`
	URL        string `json:"url"`
	IsFeatured bool   `json:"is_featured"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
