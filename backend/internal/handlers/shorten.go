package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"shortly/internal/metrics"
	"shortly/internal/middlewares"
	"shortly/internal/models"
	"shortly/internal/utils"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

// process URL shortening requests
func HandleShorten(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	metrics.TotalRequests.WithLabelValues(r.Method, r.URL.Path).Inc()

	var req struct {
		URL   string `json:"url"`
		Alias string `json:"alias,omitempty"`
	}

	// retrieve user ID from context
	userID, _ := r.Context().Value(middlewares.UserIDKey{}).(uint)

	log.WithFields(log.Fields{
		"method": r.Method,
		"url":    r.URL.Path,
		"userID": userID,
		"remote": r.RemoteAddr,
	}).Info("URL shortening request received")

	// decode json request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.WithFields(log.Fields{
			"error":  err.Error(),
			"remote": r.RemoteAddr,
		}).Warn("Invalid input for URL shortening")
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// do not allow domain url to be itself
	reqURL := strings.TrimSuffix(req.URL, "/")
	baseDomain := "https://shrink.lol"
	if strings.HasPrefix(reqURL, baseDomain) {
		log.WithFields(log.Fields{
			"attempted_url": req.URL,
			"userID":        userID,
			"remote":        r.RemoteAddr,
		}).Warn("Attempted to shorten the base domain URL")
		http.Error(w, "That URL has been blocked", http.StatusBadRequest)
		return
	}

	// generate alias if alias not already provided
	alias := req.Alias
	if alias == "" || userID == 0 {
		alias = utils.GenerateHash()
	} else {
		if existingURL, _ := utils.FindURL(alias); existingURL != nil {
			log.WithFields(log.Fields{
				"alias":  alias,
				"userID": userID,
				"remote": r.RemoteAddr,
			}).Warn("Alias already exists")
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

	base := os.Getenv("BASE_URL")
	if base == "" {
		base = "http://" + r.Host
	}

	shortened := base + "/s/" + alias

	url := models.URL{
		Alias:     alias,
		URL:       req.URL,
		Link:      shortened,
		CreatedAt: time.Now(),
		ExpiresAt: expiresAt,
		UserID:    userID,
	}

	// save url to database
	if err := utils.SaveURL(&url); err != nil {
		log.WithFields(log.Fields{
			"alias":  alias,
			"userID": userID,
			"error":  err.Error(),
			"remote": r.RemoteAddr,
		}).Error("Failed to save URL to the database")
		http.Error(w, "Could not save the URL", http.StatusInternalServerError)
		metrics.TotalErrors.WithLabelValues(r.Method, r.URL.Path, http.StatusText(http.StatusInternalServerError)).Inc()
		return
	}

	res := map[string]string{
		"shortened_url": shortened,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)

	metrics.RequestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(time.Since(start).Seconds())

	log.WithFields(log.Fields{
		"shortened_url": shortened,
		"alias":         alias,
		"userID":        userID,
		"remote":        r.RemoteAddr,
	}).Info("URL shortened successfully")
}
