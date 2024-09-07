package utils

import (
	"net/http"
	"shortly/internal/metrics"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// validates user name and password and logs user in
func Authenticate(w http.ResponseWriter, username, password string) error {
	log.WithFields(log.Fields{
		"username": username,
	}).Info("Authentication attempt")

	// get user from db using username
	user, err := FindUserByUsername(username)
	if err != nil {
		log.WithFields(log.Fields{
			"username": username,
			"error":    err.Error(),
		}).Warn("User not found during login attempt")
		return err
	}

	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		log.WithFields(log.Fields{
			"username": username,
		}).Warn("Invalid password during login attempt")
		return err
	}

	// generate JWT token
	token, err := GenerateJWT(user.ID)
	if err != nil {
		log.WithFields(log.Fields{
			"username": username,
			"error":    err.Error(),
		}).Error("Failed to generate JWT during login")
		metrics.TotalErrors.WithLabelValues("POST", "/login", "TokenGenerationFailed").Inc()
		return err
	}

	// set JWT token in HTTP-only cookie
	SetCookie(w, token)

	return nil
}
