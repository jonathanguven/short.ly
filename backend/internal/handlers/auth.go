package handlers

import (
	"encoding/json"
	"net/http"
	"shortly/internal/models"
	"shortly/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

// handle user authentication
func HandleLogin(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// find user by username
	user, err := models.FindUserByUsername(req.Username)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// check password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// generate login token
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		http.Error(w, "Failure generating token", http.StatusInternalServerError)
		return
	}

	utils.SetCookie(w, token)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login successful"))
}
