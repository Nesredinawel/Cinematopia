package handlers

import (
	"cinema-backend/services"
	"encoding/json"
	"log"
	"net/http"
)

// StarInput represents the input structure for a star
type StarInput struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Bio       string `json:"bio"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// CreateStarHandler handles the creation of a new star
func CreateStarHandler(w http.ResponseWriter, r *http.Request) {
	var input StarInput

	// Decode the request body into the StarInput struct
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Define the GraphQL mutation for inserting a star
	query := `
        mutation ($name: String!, $bio: String!) {
            insert_stars(objects: {
                name: $name,
                bio: $bio
            }) {
                affected_rows
                returning {
                    id
                    name
                    bio
                    created_at
                    updated_at
                }
            }
        }
    `

	// Prepare the variables for the GraphQL mutation
	variables := map[string]interface{}{
		"name": input.Name,
		"bio":  input.Bio,
	}

	// Call the QueryHasura function to execute the mutation
	resp, err := services.QueryHasura(query, variables)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Printf("Raw Hasura Response: %s", resp.Body)
		http.Error(w, "Error inserting star", http.StatusInternalServerError)
		return
	}

	// Decode the response body
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println("Decoding error:", err)
		http.Error(w, "Error decoding response", http.StatusInternalServerError)
		return
	}

	// Return the inserted star details
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// GetStarsHandler handles fetching all stars
func GetStarsHandler(w http.ResponseWriter, r *http.Request) {
	// Define the GraphQL query for retrieving stars
	query := `
		query {
			stars {
				id
				name
				bio
				created_at
				updated_at
			}
		}
	`

	// Call the QueryHasura function to execute the query
	resp, err := services.QueryHasura(query, nil)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Printf("Raw Hasura Response: %s", resp.Body)
		http.Error(w, "Error fetching stars", http.StatusInternalServerError)
		return
	}

	// Decode the response body
	var result struct {
		Data struct {
			Stars []StarInput `json:"stars"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println("Decoding error:", err)
		http.Error(w, "Error decoding response", http.StatusInternalServerError)
		return
	}

	// Return the stars list
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result.Data.Stars)
}
