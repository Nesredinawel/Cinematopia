package handlers

import (
	"cinema-backend/services"
	"encoding/json"
	"log"
	"net/http"
)

// Image represents an image object
type Image struct {
	ID         int    `json:"id"`
	MovieID    int    `json:"movie_id"`
	URL        string `json:"url"`
	IsFeatured bool   `json:"is_featured"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

func CreateImageHandler(w http.ResponseWriter, r *http.Request) {
	// Define input structure
	var input struct {
		MovieID    int    `json:"movie_id"`
		URL        string `json:"url"`
		IsFeatured bool   `json:"is_featured"`
	}

	// Decode the JSON request body into input struct
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Define the GraphQL mutation for inserting a new image
	query := `
		mutation ($movie_id: Int!, $url: String!, $is_featured: Boolean!) {
			insert_images(objects: {
				movie_id: $movie_id,
				url: $url,
				is_featured: $is_featured
			}) {
				affected_rows
				returning {
					id
					movie_id
					url
					is_featured
					created_at
					updated_at
				}
			}
		}
	`

	// Prepare the variables for the GraphQL query
	variables := map[string]interface{}{
		"movie_id":    input.MovieID,
		"url":         input.URL,
		"is_featured": input.IsFeatured,
	}

	// Call the QueryHasura function to execute the mutation
	resp, err := services.QueryHasura(query, variables)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Printf("Error inserting image: %s", err)
		http.Error(w, "Error inserting image", http.StatusInternalServerError)
		return
	}

	// Decode the response body
	var result struct {
		Data struct {
			InsertImages struct {
				AffectedRows int     `json:"affected_rows"`
				Returning    []Image `json:"returning"`
			} `json:"insert_images"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println("Decoding error:", err)
		http.Error(w, "Error decoding response", http.StatusInternalServerError)
		return
	}

	// Log affected rows to check if the insertion was successful
	log.Printf("Affected Rows: %d", result.Data.InsertImages.AffectedRows)

	// Check if the 'returning' slice has at least one element
	if result.Data.InsertImages.AffectedRows == 0 || len(result.Data.InsertImages.Returning) == 0 {
		log.Println("No image returned from Hasura mutation")
		http.Error(w, "No image returned", http.StatusInternalServerError)
		return
	}

	// Return the newly created image
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result.Data.InsertImages.Returning[0])
}

// GetImagesHandler handles fetching all images
func GetImagesHandler(w http.ResponseWriter, r *http.Request) {
	// Define the GraphQL query for retrieving images
	query := `
		query {
			images {
				id
				movie_id
				url
				is_featured
				created_at
				updated_at
			}
		}
	`

	// Call the QueryHasura function to execute the query
	resp, err := services.QueryHasura(query, nil)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Printf("Error fetching images: %s", err)
		http.Error(w, "Error fetching images", http.StatusInternalServerError)
		return
	}

	// Decode the response body
	var result struct {
		Data struct {
			Images []Image `json:"images"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println("Decoding error:", err)
		http.Error(w, "Error decoding response", http.StatusInternalServerError)
		return
	}

	// Return the list of images
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result.Data.Images)
}
