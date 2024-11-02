package handlers

import (
	"cinema-backend/services"
	"encoding/json"
	"log"
	"net/http"
)

// MovieStarInput represents the input structure for a movie-star relationship
type MovieStarInput struct {
	MovieID int `json:"movie_id"`
	StarID  int `json:"star_id"`
}

// CreateMovieStarHandler handles the creation of a new movie-star relationship
func CreateMovieStarHandler(w http.ResponseWriter, r *http.Request) {
	var input MovieStarInput

	// Decode the request body into the MovieStarInput struct
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Define the GraphQL mutation for inserting a movie-star relationship
	query := `
		mutation ($movie_id: Int!, $star_id: Int!) {
			insert_movie_stars(objects: {
				movie_id: $movie_id,
				star_id: $star_id
			}) {
				affected_rows
				returning {
					movie_id
					star_id
				}
			}
		}
	`

	// Prepare the variables for the GraphQL mutation
	variables := map[string]interface{}{
		"movie_id": input.MovieID,
		"star_id":  input.StarID,
	}

	// Call the QueryHasura function to execute the mutation
	resp, err := services.QueryHasura(query, variables)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Printf("Raw Hasura Response: %s", resp.Body)
		http.Error(w, "Error inserting movie-star relationship", http.StatusInternalServerError)
		return
	}

	// Decode the response body
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println("Decoding error:", err)
		http.Error(w, "Error decoding response", http.StatusInternalServerError)
		return
	}

	// Return the inserted movie-star relationship details
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// GetMovieStarsHandler handles fetching all movie-star relationships
func GetMovieStarsHandler(w http.ResponseWriter, r *http.Request) {
	// Define the GraphQL query for retrieving movie-stars
	query := `
		query {
			movie_stars {
				movie_id
				star_id
			}
		}
	`

	// Call the QueryHasura function to execute the query
	resp, err := services.QueryHasura(query, nil)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Printf("Raw Hasura Response: %s", resp.Body)
		http.Error(w, "Error fetching movie-stars", http.StatusInternalServerError)
		return
	}

	// Decode the response body
	var result struct {
		Data struct {
			MovieStars []MovieStarInput `json:"movie_stars"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println("Decoding error:", err)
		http.Error(w, "Error decoding response", http.StatusInternalServerError)
		return
	}

	// Return the movie-star relationships list
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result.Data.MovieStars)
}
