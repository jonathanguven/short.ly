package handlers

import (
	"encoding/json"
	"net/http"
	"shortly/internal/metrics"
	"shortly/internal/utils"
	"time"

	"shortly/internal/middlewares"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// update existing url alias
func HandleUpdateURL(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	metrics.TotalRequests.WithLabelValues(r.Method, r.URL.Path).Inc()

	var req struct {
		NewAlias string `json:"new_alias"`
	}

	vars := mux.Vars(r)
	alias := vars["alias"]

	log.WithFields(log.Fields{
		"method": r.Method,
		"url":    r.URL.Path,
		"alias":  alias,
		"remote": r.RemoteAddr,
	}).Info("Update URL alias request received")

	// make sure alias exists in the first place
	url, err := utils.FindURL(alias)
	if err != nil {
		log.WithFields(log.Fields{
			"alias": alias,
			"error": err.Error(),
		}).Warn("URL not found")
		metrics.TotalErrors.WithLabelValues(r.Method, r.URL.Path, "URLNotFound").Inc()
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	// get userID from context
	userID, ok := r.Context().Value(middlewares.UserIDKey{}).(uint)
	if !ok {
		log.WithFields(log.Fields{
			"remote": r.RemoteAddr,
		}).Warn("Unauthenticated user")
		metrics.TotalErrors.WithLabelValues(r.Method, r.URL.Path, "Unauthenticated").Inc()
		http.Error(w, "Unauthenticated", http.StatusUnauthorized)
		return
	}

	// make sure user owns the shortened URL
	if err := utils.VerifyUser(userID, url); err != nil {
		log.WithFields(log.Fields{
			"alias":  alias,
			"remote": r.RemoteAddr,
		}).Warn("Unauthorized attempt to modify URL")
		metrics.TotalErrors.WithLabelValues(r.Method, r.URL.Path, "Unauthorized").Inc()
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// parse req body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.WithFields(log.Fields{
			"error":  err.Error(),
			"remote": r.RemoteAddr,
		}).Warn("Invalid input for alias update")
		metrics.TotalErrors.WithLabelValues(r.Method, r.URL.Path, http.StatusText(http.StatusBadRequest)).Inc()
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// update alias
	url.Alias = req.NewAlias
	if err := utils.SaveURL(url); err != nil {
		log.WithFields(log.Fields{
			"alias":  alias,
			"error":  err.Error(),
			"remote": r.RemoteAddr,
		}).Error("Failed to update URL alias in database")
		metrics.TotalErrors.WithLabelValues(r.Method, r.URL.Path, "DatabaseSaveFailed").Inc()
		http.Error(w, "Failed to update URL alias", http.StatusInternalServerError)
		return
	}

	log.WithFields(log.Fields{
		"new_alias": url.Alias,
		"userID":    url.UserID,
		"remote":    r.RemoteAddr,
	}).Info("URL alias updated successfully")

	metrics.RequestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(time.Since(start).Seconds())
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("URL alias updated successfully"))
}
