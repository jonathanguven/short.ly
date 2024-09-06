package handlers

import (
	"encoding/json"
	"net/http"
	"shortly/internal/database"
	"shortly/internal/models"
)

// returns all shortened URLs from the database
func findAllURLs() ([]models.URL, error) {
	var urls []models.URL
	result := database.DB.Find(&urls)
	return urls, result.Error
}

// return all created shortened URLs
func HandleListURLs(w http.ResponseWriter, r *http.Request) {
	urls, err := findAllURLs()

	if err != nil {
		http.Error(w, "Failed to retrieve URLs", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(urls)
}
