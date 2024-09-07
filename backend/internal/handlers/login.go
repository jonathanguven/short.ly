package handlers

import (
	"encoding/json"
	"net/http"
	"shortly/internal/metrics"
	"shortly/internal/utils"
	"time"

	log "github.com/sirupsen/logrus"
)

// handle user authentication
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	metrics.TotalRequests.WithLabelValues(r.Method, r.URL.Path).Inc()

	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	log.WithFields(log.Fields{
		"method": r.Method,
		"url":    r.URL.Path,
		"remote": r.RemoteAddr,
	}).Info("Login request received")

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.WithFields(log.Fields{
			"error":  err.Error(),
			"remote": r.RemoteAddr,
		}).Warn("Invalid input during login attempt")
		metrics.TotalErrors.WithLabelValues(r.Method, r.URL.Path, http.StatusText(http.StatusBadRequest)).Inc()
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// authenticate user and log them in
	if err := utils.Authenticate(w, req.Username, req.Password); err != nil {
		log.WithFields(log.Fields{
			"username": req.Username,
			"remote":   r.RemoteAddr,
		}).Warn("Invalid username or password")
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	log.WithFields(log.Fields{
		"username": req.Username,
		"remote":   r.RemoteAddr,
	}).Info("User logged in successfully")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login successful"))

	metrics.RequestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(time.Since(start).Seconds())

}
