package models

type Bookmark struct {
	UserID    int    `json:"user_id"`
	MovieID   int    `json:"movie_id"`
	CreatedAt string `json:"created_at"`
}
