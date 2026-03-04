package handler

import (
	"encoding/json"
	"net/http"
)

// HealthResponse represents the response for the health check endpoint.
type HealthResponse struct {
	Status string `json:"status"`
}

// Health handles GET /health requests and returns 200 OK.
func Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(HealthResponse{Status: "ok"})
}
