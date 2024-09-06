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
		Alias     string `json:"alias,omitempty"`
		ExpiresAt int    `json:"expires_in,omitempty"`
	}

	// decode json request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// generate alias if alias not already provided
	alias := req.Alias
	if alias == "" {
		alias = utils.GenerateHash()
	}

	// alias already exists in database
	if existingURL, _ := models.FindURL(alias); existingURL != nil {
		http.Error(w, "Alias already exists, please choose another one", http.StatusBadRequest)
		return
	}

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

	// save url to database
	if err := models.SaveURL(&url); err != nil {
		http.Error(w, "Could not save the URL", http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"shortened_url": "http://localhost:8000/s/" + alias,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
