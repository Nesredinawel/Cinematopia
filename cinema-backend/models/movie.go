package models

// Movie represents the structure for movie data
type Movie struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Duration    int     `json:"duration"`
	ReleaseYear int     `json:"release_year"`
	Genre       string  `json:"genre"`
	Rating      float64 `json:"rating"`
	DirectorID  int     `json:"director_id"`
}
