package utils

import (
	"encoding/json"
	"net/http"
)

// RespondWithError sends an error response in JSON format
func RespondWithError(w http.ResponseWriter, statusCode int, message string) {
	RespondWithJSON(w, statusCode, map[string]string{"error": message})
}

// RespondWithJSON sends a JSON response with the provided data
func RespondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	response, _ := json.Marshal(data)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}
