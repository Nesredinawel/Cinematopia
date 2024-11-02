package handlers

import (
	"cinema-backend/services"
	"encoding/json"
	"log"
	"net/http"
)

// MovieInput represents the input structure for a movie
type MovieInput struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Duration    int     `json:"duration"`
	ReleaseYear int     `json:"release_year"`
	Genre       string  `json:"genre"`
	Rating      float64 `json:"rating"` // Use float64 to ensure compatibility with Hasura's float8
	DirectorID  int     `json:"director_id"`
}

// CreateMovieHandler handles the creation of a new movie
func CreateMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input MovieInput

	// Decode the request body into the MovieInput struct
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Define the GraphQL mutation for inserting a movie
	query := `
		mutation ($title: String!, $description: String!, $duration: Int!, $release_year: Int!, $genre: String!, $rating: float8!, $director_id: Int!) {
			insert_movies(objects: {
				title: $title,
				description: $description,
				duration: $duration,
				release_year: $release_year,
				genre: $genre,
				rating: $rating,
				director_id: $director_id
			}) {
				affected_rows
				returning {
					id
					title
					description
					duration
					release_year
					genre
					rating
					director_id
					created_at
					updated_at
				}
			}
		}
	`

	// Prepare the variables for the GraphQL mutation
	variables := map[string]interface{}{
		"title":        input.Title,
		"description":  input.Description,
		"duration":     input.Duration,
		"release_year": input.ReleaseYear,
		"genre":        input.Genre,
		"rating":       input.Rating, // Ensure this is float64
		"director_id":  input.DirectorID,
	}

	// Call the QueryHasura function to execute the mutation
	resp, err := services.QueryHasura(query, variables)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Printf("Raw Hasura Response: %s", resp.Body)
		http.Error(w, "Error inserting movie", http.StatusInternalServerError)
		return
	}

	// Decode the response body
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println("Decoding error:", err)
		http.Error(w, "Error decoding response", http.StatusInternalServerError)
		return
	}

	// Return the inserted movie details
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// GetMoviesHandler handles fetching all movies
func GetMoviesHandler(w http.ResponseWriter, r *http.Request) {
	// Define the GraphQL query for retrieving movies
	query := `
		query {
			movies {
				id
				title
				description
				duration
				release_year
				genre
				rating
				director_id
				created_at
				updated_at
			}
		}
	`

	// Call the QueryHasura function to execute the query
	resp, err := services.QueryHasura(query, nil)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Printf("Raw Hasura Response: %s", resp.Body)
		http.Error(w, "Error fetching movies", http.StatusInternalServerError)
		return
	}

	// Decode the response body
	var result struct {
		Data struct {
			ID     int          `json:"id"`
			Movies []MovieInput `json:"movies"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println("Decoding error:", err)
		http.Error(w, "Error decoding response", http.StatusInternalServerError)
		return
	}

	// Return the movies list
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result.Data.Movies)
}
