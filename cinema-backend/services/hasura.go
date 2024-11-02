package services

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

func QueryHasura(query string, variables map[string]interface{}) (*http.Response, error) {
	url := os.Getenv("HASURA_GRAPHQL_ENDPOINT")
	body := map[string]interface{}{
		"query":     query,
		"variables": variables,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Hasura-Admin-Secret", os.Getenv("HASURA_GRAPHQL_ADMIN_SECRET")) // If using admin secret

	client := &http.Client{}
	return client.Do(req)
}
