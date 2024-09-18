package handlers

import (
	"net/http"
	"os"
	"shortly/internal/metrics"
	"time"

	log "github.com/sirupsen/logrus"
)

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	metrics.TotalRequests.WithLabelValues(r.Method, r.URL.Path).Inc()

	log.WithFields(log.Fields{
		"method": r.Method,
		"url":    r.URL.Path,
		"remote": r.RemoteAddr,
	}).Info("Logout request received")

	domain := os.Getenv("API_DOMAIN")
	secure := os.Getenv("ENVIRONMENT") == "production"
	sameSite := http.SameSiteNoneMode

	if os.Getenv("ENVIRONMENT") != "production" {
		domain = ""
		secure = false
		sameSite = http.SameSiteLaxMode
	}

	cookie := &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		Domain:   domain,
		Expires:  time.Now().Add(time.Hour * 24 * 7),
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   secure,
		SameSite: sameSite,
	}

	http.SetCookie(w, cookie)

	log.WithFields(log.Fields{
		"remote": r.RemoteAddr,
	}).Info("User logged out successfully")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logout successful"))

	metrics.RequestDuration.WithLabelValues(r.Method, r.URL.Path).Observe(time.Since(start).Seconds())
}
