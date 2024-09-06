package handlers

import (
	"encoding/json"
	"net/http"
	"shortly/internal/utils"
)

// return all created shortened URLs
func HandleListURLs(w http.ResponseWriter, r *http.Request) {
	urls, err := utils.FindAllURLs()

	if err != nil {
		http.Error(w, "Failed to retrieve URLs", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(urls)
}
