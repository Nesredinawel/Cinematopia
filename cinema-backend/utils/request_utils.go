package utils

import (
	"encoding/json"
	"net/http"
)

// ParseJSONRequest parses the JSON body of a request into the specified struct.
func ParseJSONRequest(r *http.Request, target interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(target)
	if err != nil {
		return err
	}
	return nil
}

// SendJSONResponse sends a JSON response with the specified data and status code.
func SendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

// SendErrorResponse sends an error response with a given status code and message.
func SendErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	SendJSONResponse(w, statusCode, map[string]string{"error": message})
}

// IsValidRequestMethod checks if the HTTP request method matches the expected method.
func IsValidRequestMethod(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		SendErrorResponse(w, http.StatusMethodNotAllowed, "Invalid request method")
		return false
	}
	return true
}

// CheckAuthHeader checks if the Authorization header is present and returns its value.
func CheckAuthHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", http.ErrNoCookie // Using ErrNoCookie to signify missing auth header
	}
	return authHeader, nil
}
