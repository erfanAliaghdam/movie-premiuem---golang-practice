package handler

import (
	"encoding/json"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// Define a response struct or map
	response := map[string]string{
		"message": "Welcome to Movie Premium!",
		"status":  "success",
	}

	// Set headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode and write the JSON response
	json.NewEncoder(w).Encode(response)
}
