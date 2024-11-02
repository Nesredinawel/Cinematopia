package handlers

import (
	"cinema-backend/services"
	"encoding/json"
	"log"
	"net/http"
)

// ScheduleInput represents the input structure for a schedule
type ScheduleInput struct {
	ID             int    `json:"id"`
	CinemaLocation string `json:"cinema_location"`
	AvailableSeats int    `json:"available_seats"`
	Price          int    `json:"Price"`
	MovieID        int    `json:"movie_id"`
	ShowTime       string `json:"show_time"` // Use string for the show_time in ISO 8601 format (e.g., "2024-10-22T19:30:00Z")
}

// CreateScheduleHandler handles the creation of a new schedule
func CreateScheduleHandler(w http.ResponseWriter, r *http.Request) {
	var input ScheduleInput

	// Decode the request body into the ScheduleInput struct
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Define the GraphQL mutation for inserting a schedule
	query := `
		mutation ($cinema_location: String!, $available_seats: Int!, $Price: Int!, $movie_id: Int!, $show_time: timestamp!) {
			insert_schedules(objects: {
				cinema_location: $cinema_location,
				available_seats: $available_seats,
				Price: $Price,
				movie_id: $movie_id,
				show_time: $show_time
			}) {
				affected_rows
				returning {
					id
					cinema_location
					available_seats
					Price
					movie_id
					show_time
					created_at
					updated_at
				}
			}
		}
	`

	// Prepare the variables for the GraphQL mutation
	variables := map[string]interface{}{
		"cinema_location": input.CinemaLocation,
		"available_seats": input.AvailableSeats,
		"movie_id":        input.MovieID,
		"show_time":       input.ShowTime,
		"Price":           input.Price,
	}

	// Call the QueryHasura function to execute the mutation
	resp, err := services.QueryHasura(query, variables)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Printf("Raw Hasura Response: %s", resp.Body)
		http.Error(w, "Error inserting schedule", http.StatusInternalServerError)
		return
	}

	// Decode the response body
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println("Decoding error:", err)
		http.Error(w, "Error decoding response", http.StatusInternalServerError)
		return
	}

	// Return the inserted schedule details
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// GetSchedulesHandler handles fetching all schedules
func GetSchedulesHandler(w http.ResponseWriter, r *http.Request) {
	// Define the GraphQL query for retrieving schedules
	query := `
		query {
			schedules {
				id
				cinema_location
				available_seats
				movie_id
				show_time
				Price
				created_at
				updated_at
			}
		}
	`

	// Call the QueryHasura function to execute the query
	resp, err := services.QueryHasura(query, nil)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Printf("Raw Hasura Response: %s", resp.Body)
		http.Error(w, "Error fetching schedules", http.StatusInternalServerError)
		return
	}

	// Decode the response body
	var result struct {
		Data struct {
			Schedules []ScheduleInput `json:"schedules"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println("Decoding error:", err)
		http.Error(w, "Error decoding response", http.StatusInternalServerError)
		return
	}

	// Return the schedules list
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result.Data.Schedules)
}
