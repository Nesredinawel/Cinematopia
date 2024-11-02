package handlers

import (
	"cinema-backend/services" // Import your services package
	"encoding/json"
	"net/http"
)

// GetDirectorsHandler handles the retrieval of directors from the database
func GetDirectorsHandler(w http.ResponseWriter, r *http.Request) {
	// Define the GraphQL query for fetching directors
	query := `
		query {
			directors {
				id
				name
			}
		}
	`

	// Call the QueryHasura function to execute the query
	resp, err := services.QueryHasura(query, nil)
	if err != nil || resp.StatusCode != http.StatusOK {
		http.Error(w, "Error fetching directors", http.StatusInternalServerError)
		return
	}

	// Decode the response from Hasura
	var response struct {
		Data struct {
			Directors []struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"directors"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		http.Error(w, "Error decoding response", http.StatusInternalServerError)
		return
	}

	// Respond with the list of directors
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response.Data.Directors)

}

// CreateDirectorHandler handles the creation of a new director
func CreateDirectorHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string `json:"name"`
		Bio  string `json:"bio"` // Assuming you might want to include a bio
	}

	// Decode the JSON request body into input struct
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil || input.Name == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Define the GraphQL mutation for inserting a new director
	query := `
		mutation ($name: String!, $bio: String) {
			insert_directors_one(object: {
				name: $name,
				bio: $bio
			}) {
				id
				name
				bio
				created_at
				updated_at
			}
		}
	`

	// Prepare the variables for the GraphQL query
	variables := map[string]interface{}{
		"name": input.Name,
		"bio":  input.Bio,
	}

	// Call the QueryHasura function to execute the mutation
	resp, err := services.QueryHasura(query, variables)
	if err != nil || resp.StatusCode != http.StatusOK {
		http.Error(w, "Error inserting director", http.StatusInternalServerError)
		return
	}

	// Decode the response from Hasura
	var response struct {
		Data struct {
			InsertDirector struct {
				ID        int    `json:"id"`
				Name      string `json:"name"`
				Bio       string `json:"bio"`
				CreatedAt string `json:"created_at"`
				UpdatedAt string `json:"updated_at"`
			} `json:"insert_directors_one"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		http.Error(w, "Error decoding response", http.StatusInternalServerError)
		return
	}

	// Respond with the created director information
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response.Data.InsertDirector)
}
