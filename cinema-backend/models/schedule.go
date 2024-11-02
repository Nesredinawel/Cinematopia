package models

type Schedule struct {
	ID             int    `json:"id"`
	MovieID        int    `json:"movie_id"`
	CinemaLocation string `json:"cinema_location"`
	ShowTime       string `json:"show_time"`
	AvailableSeats int    `json:"available_seats"`
}
