package handlers

import (
	"cinema-backend/services"
	"cinema-backend/utils"
	"encoding/json"
	"net/http"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Decode the JSON request body into input struct
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Hash the password
	passwordHash, err := utils.HashPassword(input.Password)
	if err != nil {
		http.Error(w, "Error creating password hash", http.StatusInternalServerError)
		return
	}

	// Define the GraphQL mutation for inserting a new user
	query := `
		mutation ($name: String!, $email: String!, $password_hash: String!) {
			insert_users_one(object: {
				name: $name,
				email: $email,
				password_hash: $password_hash
			}) {
				id
				email
			}
		}
	`

	// Prepare the variables for the GraphQL query
	variables := map[string]interface{}{
		"name":          input.Name,
		"email":         input.Email,
		"password_hash": passwordHash,
	}

	// Call the QueryHasura function to execute the mutation
	resp, err := services.QueryHasura(query, variables)
	if err != nil || resp.StatusCode != http.StatusOK {
		http.Error(w, "Error inserting user", http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Decode the JSON request body into input struct
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Define the GraphQL query for retrieving the user by email
	query := `
		query ($email: String!) {
			users(where: {email: {_eq: $email}}) {
				id
				password_hash
			}
		}
	`

	// Prepare the variables for the GraphQL query
	variables := map[string]interface{}{
		"email": input.Email,
	}

	// Call the QueryHasura function to execute the query
	resp, err := services.QueryHasura(query, variables)
	if err != nil || resp.StatusCode != http.StatusOK {
		http.Error(w, "Error querying user", http.StatusInternalServerError)
		return
	}

	// Parse the response to get user data
	var data struct {
		Data struct {
			Users []struct {
				ID           int    `json:"id"`
				PasswordHash string `json:"password_hash"`
			} `json:"users"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		http.Error(w, "Error decoding response", http.StatusInternalServerError)
		return
	}

	// Check if the user exists
	if len(data.Data.Users) == 0 {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	// Verify the password
	user := data.Data.Users[0]
	if !utils.CheckPasswordHash(input.Password, user.PasswordHash) {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	// Respond with the token
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// GetUsersHandler retrieves all users from the database
func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Define the GraphQL query for retrieving all users
	query := `
		query {
			users {
				id
				name
				email
			}
		}
	`

	// Call the QueryHasura function to execute the query
	resp, err := services.QueryHasura(query, nil)
	if err != nil || resp.StatusCode != http.StatusOK {
		http.Error(w, "Error querying users", http.StatusInternalServerError)
		return
	}

	// Parse the response to get user data
	var data struct {
		Data struct {
			Users []struct {
				ID    int    `json:"id"`
				Name  string `json:"name"`
				Email string `json:"email"`
			} `json:"users"`
		} `json:"data"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		http.Error(w, "Error decoding response", http.StatusInternalServerError)
		return
	}

	// Respond with the user data
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.Data.Users)
}
