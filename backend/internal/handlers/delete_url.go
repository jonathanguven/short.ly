package handlers

import (
	"encoding/json"
	"net/http"
	"shortly/internal/metrics"
	"shortly/internal/utils"
	"strconv"
	"time"

	"shortly/internal/middlewares"

	log "github.com/sirupsen/logrus"
)

type DeleteURLsRequest struct {
	URLs []string `json:"urls"`
}

// handle deleting a URL
func HandleDeleteURL(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	metrics.TotalRequests.WithLabelValues(r.Method, r.URL.Path).Inc()

	var requestBody DeleteURLsRequest
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		log.WithFields(log.Fields{
			"error": err.Error(),
		}).Error("Failed to decode request body")
		metrics.TotalErrors.WithLabelValues(r.Method, r.URL.Path, "BadRequest").Inc()
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userID, ok := r.Context().Value(middlewares.UserIDKey{}).(uint)
	if !ok {
		log.WithFields(log.Fields{
			"remote": r.RemoteAddr,
		}).Warn("Unauthenticated user")
		metrics.TotalErrors.WithLabelValues(r.Method, r.URL.Path, "Unauthenticated").Inc()
		http.Error(w, "Unauthenticated", http.StatusUnauthorized)
		return
	}

	successfulDeletions := 0
	var failedAliases []string

	for _, alias := range requestBody.URLs {
		log.WithFields(log.Fields{
			"method": r.Method,
			"url":    r.URL.Path,
			"alias":  alias,
			"remote": r.RemoteAddr,
		}).Info("Delete URL request received")

		url, err := utils.FindURL(alias)
		if err != nil {
			log.WithFields(log.Fields{
				"alias": alias,
				"error": err.Error(),
			}).Warn("URL not found")
			metrics.TotalErrors.WithLabelValues(r.Method, r.URL.Path, "URLNotFound").Inc()
			failedAliases = append(failedAliases, alias)
			continue
		}

		if err := utils.VerifyUser(userID, url); err != nil {
			log.WithFields(log.Fields{
				"alias":  alias,
				"remote": r.RemoteAddr,
			}).Warn("Unauthorized attempt to delete URL")
			metrics.TotalErrors.WithLabelValues(r.Method, r.URL.Path, "Unauthorized").Inc()
			failedAliases = append(failedAliases, alias)
			continue
		}

		if err := utils.DeleteURL(url.ID); err != nil {
			log.WithFields(log.Fields{
				"alias":  alias,
				"error":  err.Error(),
				"remote": r.RemoteAddr,
			}).Error("Failed to delete URL from database")
			metrics.TotalErrors.WithLabelValues(r.Method, r.URL.Path, "DatabaseDeleteFailed").Inc()
			failedAliases = append(failedAliases, alias)
			continue
		}

		log.WithFields(log.Fields{
			"alias":  alias,
			"userID": url.UserID,
			"remote": r.RemoteAddr,
		}).Info("URL deleted successfully")
		successfulDeletions++
	}

	metrics.RequestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(time.Since(start).Seconds())

	if len(failedAliases) > 0 {
		w.WriteHeader(http.StatusPartialContent)
		w.Write([]byte("Failed to delete the following URLs: " + jsonStringify(failedAliases)))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strconv.Itoa(successfulDeletions) + " URLs deleted successfully"))
}

func jsonStringify(data []string) string {
	result, _ := json.Marshal(data)
	return string(result)
}
