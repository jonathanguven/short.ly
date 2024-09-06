package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"shortly/internal/middlewares"
	"shortly/internal/models"
	"shortly/internal/utils"
	"time"
)

// process URL shortening requests
func HandleShorten(w http.ResponseWriter, r *http.Request) {
	var req struct {
		URL   string `json:"url"`
		Alias string `json:"alias,omitempty"`
	}

	// retrieve user ID from context
	userID, _ := r.Context().Value(middlewares.UserIDKey{}).(uint)

	// decode json request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// generate alias if alias not already provided
	alias := req.Alias
	if alias == "" || userID == 0 {
		alias = utils.GenerateHash()
	} else {
		if existingURL, _ := utils.FindURL(alias); existingURL != nil {
			http.Error(w, "Alias already exists, please choose another one", http.StatusBadRequest)
			return
		}
	}

	var expiresAt *time.Time
	if userID == 0 { // set a default expiration for guest user (7 days)
		expiration := time.Now().Add(7 * 24 * time.Hour)
		expiresAt = &expiration
	} else { // logged in user is creating the shortened URL
		expiresAt = nil
	}

	// alias already exists in database
	if existingURL, _ := utils.FindURL(alias); existingURL != nil {
		http.Error(w, "Alias already exists, please choose another one", http.StatusBadRequest)
		return
	}

	url := models.URL{
		Alias:     alias,
		URL:       req.URL,
		CreatedAt: time.Now(),
		ExpiresAt: expiresAt,
	}

	// save url to database
	if err := utils.SaveURL(&url); err != nil {
		http.Error(w, "Could not save the URL", http.StatusInternalServerError)
		return
	}

	base := os.Getenv("BASE_URL")
	if base == "" {
		base = "http://" + r.Host
	}

	res := map[string]string{
		"shortened_url": base + "/s/" + alias,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
