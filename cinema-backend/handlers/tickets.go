package handlers

import (
	"cinema-backend/services"
	"encoding/json"
	"log"
	"net/http"
)

// TicketInput represents the input structure for a ticket
type TicketInput struct {
	PaymentStatus string  `json:"payment_status"`
	Quantity      int     `json:"quantity"`
	ScheduleID    int     `json:"schedule_id"`
	UserID        int     `json:"user_id"`
	TotalPrice    float64 `json:"total_price"`
	QRCodeURL     string  `json:"qr_code_url"`
}

// CreateTicketHandler handles the creation of a new ticket
func CreateTicketHandler(w http.ResponseWriter, r *http.Request) {
	var input TicketInput

	// Decode the request body into the TicketInput struct
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Define the GraphQL mutation for inserting a ticket
	query := `
		mutation ($quantity: Int!, $schedule_id: Int!, $user_id: Int!, $total_price: numeric!, $payment_status: String!, $qr_code_url: String!) {
			insert_tickets(objects: {
				quantity: $quantity,
				schedule_id: $schedule_id,
				user_id: $user_id,
				total_price: $total_price,
				payment_status: $payment_status,
				qr_code_url: $qr_code_url
			}) {
				affected_rows
				returning {
					id
					quantity
					schedule_id
					user_id
					total_price
					payment_status
					qr_code_url
					created_at
					
				}
			}
		}
	`

	// Prepare the variables for the GraphQL mutation
	variables := map[string]interface{}{
		"quantity":       input.Quantity,
		"schedule_id":    input.ScheduleID,
		"user_id":        input.UserID,
		"total_price":    input.TotalPrice,
		"payment_status": input.PaymentStatus,
		"qr_code_url":    input.QRCodeURL,
	}

	// Call the QueryHasura function to execute the mutation
	resp, err := services.QueryHasura(query, variables)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Printf("Raw Hasura Response: %s", resp.Body)
		http.Error(w, "Error inserting ticket", http.StatusInternalServerError)
		return
	}

	// Decode the response body
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println("Decoding error:", err)
		http.Error(w, "Error decoding response", http.StatusInternalServerError)
		return
	}

	// Return the inserted ticket details
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// GetTicketsHandler handles fetching all tickets
func GetTicketsHandler(w http.ResponseWriter, r *http.Request) {
	// Define the GraphQL query for retrieving tickets
	query := `
		query {
			tickets {
				id
				quantity
				schedule_id
				user_id
				total_price
				payment_status
				qr_code_url
				created_at
				updated_at
			}
		}
	`

	// Call the QueryHasura function to execute the query
	resp, err := services.QueryHasura(query, nil)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Printf("Raw Hasura Response: %s", resp.Body)
		http.Error(w, "Error fetching tickets", http.StatusInternalServerError)
		return
	}

	// Decode the response body
	var result struct {
		Data struct {
			Tickets []TicketInput `json:"tickets"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println("Decoding error:", err)
		http.Error(w, "Error decoding response", http.StatusInternalServerError)
		return
	}

	// Return the tickets list
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result.Data.Tickets)
}
