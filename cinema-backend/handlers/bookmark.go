package handlers

import (
	"cinema-backend/services"
	"encoding/json"
	"log"
	"net/http"
)

// InsertBookmarkRequest represents the request body for inserting a bookmark
type InsertBookmarkRequest struct {
	MovieID int `json:"movie_id"`
	UserID  int `json:"user_id"`
}

// Bookmark represents a bookmark object
type Bookmark struct {
	MovieID   int    `json:"movie_id"`
	UserID    int    `json:"user_id"`
	CreatedAt string `json:"created_at"`
}

// CreateBookmarkHandler handles the insertion of a new bookmark
func CreateBookmarkHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body into the InsertBookmarkRequest struct
	var input InsertBookmarkRequest
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Define the GraphQL mutation for inserting a new bookmark
	query := `
		mutation ($movie_id: Int!, $user_id: Int!) {
			insert_bookmarks(objects: {
				movie_id: $movie_id,
				user_id: $user_id
			}) {
				affected_rows
				returning {
					movie_id
					user_id
					created_at
				}
			}
		}
	`

	// Prepare the variables for the GraphQL mutation
	variables := map[string]interface{}{
		"movie_id": input.MovieID,
		"user_id":  input.UserID,
	}

	// Call the QueryHasura function to execute the mutation
	resp, err := services.QueryHasura(query, variables)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Printf("Raw Hasura Response: %s", resp.Body)
		http.Error(w, "Error inserting bookmark", http.StatusInternalServerError)
		return
	}

	// Decode the response body to get the inserted bookmark data
	var result struct {
		Data struct {
			InsertBookmarks struct {
				Returning []Bookmark `json:"returning"`
			} `json:"insert_bookmarks"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println("Decoding error:", err)
		http.Error(w, "Error decoding response", http.StatusInternalServerError)
		return
	}

	// Return the inserted bookmark details as a JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result.Data.InsertBookmarks.Returning)
}

// GetBookmarksHandler handles retrieving the bookmarks
func GetUserBookmarksHandler(w http.ResponseWriter, r *http.Request) {
	// Define the GraphQL query for fetching bookmarks
	query := `
		query {
			bookmarks {
				movie_id
				user_id
				created_at
			}
		}
	`

	// Call the QueryHasura function to execute the query
	resp, err := services.QueryHasura(query, nil)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Printf("Raw Hasura Response: %s", resp.Body)
		http.Error(w, "Error fetching bookmarks", http.StatusInternalServerError)
		return
	}

	// Decode the response body to get the bookmark data
	var result struct {
		Data struct {
			Bookmarks []Bookmark `json:"bookmarks"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println("Decoding error:", err)
		http.Error(w, "Error decoding response", http.StatusInternalServerError)
		return
	}

	// Return the fetched bookmarks as a JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result.Data.Bookmarks)
}
