package handlers

import (
	"encoding/json"
	"net/http"
	"shortly/internal/models"
	"shortly/internal/utils"
	"time"
)

// process URL shortening requests
func HandleShorten(w http.ResponseWriter, r *http.Request) {
	var req struct {
		URL       string `json:"url"`
		ExpiresAt int    `json:"expires_in,omitempty"`
	}

	// decode json request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// generate alias
	alias := utils.GenerateHash()

	var expiresAt *time.Time
	if req.ExpiresAt > 0 {
		expiration := time.Now().Add(time.Duration(req.ExpiresAt) * 24 * time.Hour)
		expiresAt = &expiration
	}

	url := models.URL{
		Alias:     alias,
		URL:       req.URL,
		CreatedAt: time.Now(),
		ExpiresAt: expiresAt,
	}
	models.SaveURL(&url)

	res := map[string]string{
		"shortened_url": "http://localhost:8000/" + alias,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
