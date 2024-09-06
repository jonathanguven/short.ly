package handlers

import (
	"net/http"
	"shortly/internal/models"
	"time"
)

// process URL redirection requests
func HandleRedirect(w http.ResponseWriter, r *http.Request) {
	// extract the URL alias
	alias := r.URL.Path[len("/redirect/"):]

	// find alias in DB
	url, err := models.FindURL(alias)
	if err != nil || (url.ExpiresAt != nil && time.Now().After(*url.ExpiresAt)) {
		http.Error(w, "URL expired or not found", http.StatusNotFound)
		return
	}

	// redirect to original URL
	http.Redirect(w, r, url.URL, http.StatusFound)
}
