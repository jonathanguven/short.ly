package handlers

import (
	"encoding/json"
	"net/http"
	"shortly/internal/metrics"
	"shortly/internal/utils"
	"time"

	"shortly/internal/middlewares"

	log "github.com/sirupsen/logrus"
)

// return all created shortened URLs
func HandleListURLs(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	metrics.TotalRequests.WithLabelValues(r.Method, r.URL.Path).Inc()

	userID, ok := r.Context().Value(middlewares.UserIDKey{}).(uint)
	if !ok || userID == 0 {
		log.Warn("Unauthorized access attempt to List URLs")
		metrics.TotalErrors.WithLabelValues(r.Method, r.URL.Path, "Unauthorized").Inc()
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	log.WithFields(log.Fields{
		"userID": userID,
		"method": r.Method,
		"url":    r.URL.Path,
		"remote": r.RemoteAddr,
	}).Info("List URLs request received")

	urls, err := utils.FindURLsByUser(userID)
	if err != nil {
		log.WithFields(log.Fields{
			"userID": userID,
			"error":  err.Error(),
		}).Error("Failed to retrieve URLs for user")
		metrics.TotalErrors.WithLabelValues(r.Method, r.URL.Path, "DatabaseError").Inc()
		http.Error(w, "Failed to retrieve URLs", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(urls)

	log.WithFields(log.Fields{
		"userID": userID,
		"count":  len(urls),
	}).Info("Successfully retrieved user URLs")

	metrics.RequestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(time.Since(start).Seconds())
}
