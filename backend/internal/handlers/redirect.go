package handlers

import (
	"net/http"
	"shortly/internal/utils"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// process URL redirection requests
func HandleRedirect(w http.ResponseWriter, r *http.Request) {
	// extract the URL alias
	vars := mux.Vars(r)
	alias := vars["alias"]

	log.WithFields(log.Fields{
		"method": r.Method,
		"url":    r.URL.Path,
		"alias":  alias,
		"remote": r.RemoteAddr,
	}).Info("Redirection request received")

	// find alias in DB
	url, err := utils.FindURL(alias)
	if err != nil || (url.ExpiresAt != nil && time.Now().After(*url.ExpiresAt)) {
		log.WithFields(log.Fields{
			"alias":  alias,
			"remote": r.RemoteAddr,
		}).Warn("URL expired or not found")
		http.Error(w, "URL expired or not found", http.StatusNotFound)
		return
	}

	log.WithFields(log.Fields{
		"alias":  alias,
		"url":    url.URL,
		"remote": r.RemoteAddr,
	}).Info("Redirecting to original URL")

	// redirect to original URL
	http.Redirect(w, r, url.URL, http.StatusFound)
}
