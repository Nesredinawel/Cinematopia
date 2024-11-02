package handlers

import (
	"cinema-backend/services"
	"encoding/json"
	"log"
	"net/http"
)

// SeatInput represents the input structure for a seat
type SeatInput struct {
	ScheduleID  int    `json:"schedule_id"`
	SeatNumber  string `json:"seat_number"`
	Row         int    `json:"row"`
	SeatColumn  int    `json:"seat_column"`
	IsAvailable bool   `json:"is_available"`
}

// CreateSeatHandler handles the creation of a new seat
func CreateSeatHandler(w http.ResponseWriter, r *http.Request) {
	var input SeatInput

	// Decode the request body into the SeatInput struct
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Define the GraphQL mutation for inserting a seat
	query := `
		mutation ($schedule_id: Int!, $seat_number: String!, $row: Int!, $seat_column: Int!, $is_available: Boolean!) {
			insert_seats(objects: {
				schedule_id: $schedule_id,
				seat_number: $seat_number,
				row: $row,
				seat_column: $seat_column,
				is_available: $is_available
			}) {
				affected_rows
				returning {
					id
					schedule_id
					seat_number
					row
					seat_column
					is_available
					created_at
					updated_at
				}
			}
		}
	`

	// Prepare the variables for the GraphQL mutation
	variables := map[string]interface{}{
		"schedule_id":  input.ScheduleID,
		"seat_number":  input.SeatNumber,
		"row":          input.Row,
		"seat_column":  input.SeatColumn,
		"is_available": input.IsAvailable,
	}

	// Call the QueryHasura function to execute the mutation
	resp, err := services.QueryHasura(query, variables)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Printf("Raw Hasura Response: %s", resp.Body)
		http.Error(w, "Error inserting seat", http.StatusInternalServerError)
		return
	}

	// Decode the response body
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println("Decoding error:", err)
		http.Error(w, "Error decoding response", http.StatusInternalServerError)
		return
	}

	// Return the inserted seat details
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// GetSeatsHandler handles fetching all seats
func GetSeatsHandler(w http.ResponseWriter, r *http.Request) {
	// Define the GraphQL query for retrieving seats
	query := `
		query {
			seats {
				id
				schedule_id
				seat_number
				row
				seat_column
				is_available
				created_at
				updated_at
			}
		}
	`

	// Call the QueryHasura function to execute the query
	resp, err := services.QueryHasura(query, nil)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Printf("Raw Hasura Response: %s", resp.Body)
		http.Error(w, "Error fetching seats", http.StatusInternalServerError)
		return
	}

	// Decode the response body
	var result struct {
		Data struct {
			Seats []SeatInput `json:"seats"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println("Decoding error:", err)
		http.Error(w, "Error decoding response", http.StatusInternalServerError)
		return
	}

	// Return the seats list
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result.Data.Seats)
}
