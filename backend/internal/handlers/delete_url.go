package handlers

import (
	"net/http"
	"shortly/internal/metrics"
	"shortly/internal/utils"
	"time"

	"shortly/internal/middlewares"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// handle deleting a URL
func HandleDeleteURL(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	metrics.TotalRequests.WithLabelValues(r.Method, r.URL.Path).Inc()

	vars := mux.Vars(r)
	alias := vars["alias"]

	log.WithFields(log.Fields{
		"method": r.Method,
		"url":    r.URL.Path,
		"alias":  alias,
		"remote": r.RemoteAddr,
	}).Info("Delete URL request received")

	// make sure alias exists in db
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

	// make sure user owns url
	if err := utils.VerifyUser(userID, url); err != nil {
		log.WithFields(log.Fields{
			"alias":  alias,
			"remote": r.RemoteAddr,
		}).Warn("Unauthorized attempt to delete URL")
		metrics.TotalErrors.WithLabelValues(r.Method, r.URL.Path, "Unauthorized").Inc()
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// delete :P
	if err := utils.DeleteURL(url.ID); err != nil {
		log.WithFields(log.Fields{
			"alias":  alias,
			"error":  err.Error(),
			"remote": r.RemoteAddr,
		}).Error("Failed to delete URL from database")
		metrics.TotalErrors.WithLabelValues(r.Method, r.URL.Path, "DatabaseDeleteFailed").Inc()
		http.Error(w, "Failed to delete URL", http.StatusInternalServerError)
		return
	}

	log.WithFields(log.Fields{
		"alias":  alias,
		"userID": url.UserID,
		"remote": r.RemoteAddr,
	}).Info("URL deleted successfully")

	metrics.RequestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(time.Since(start).Seconds())
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("URL deleted successfully"))
}
