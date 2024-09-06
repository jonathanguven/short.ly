package handlers

import (
	"encoding/json"
	"net/http"
	"shortly/internal/models"
	"shortly/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

// creates new user then logs them in
func HandleCreateUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// parse req body
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	_, err := models.FindUserByUsername(req.Username)
	if err == nil {
		http.Error(w, "Username already exists", http.StatusBadRequest)
		return
	}

	// hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	// create user and save into db
	user := models.User{
		Username:     req.Username,
		PasswordHash: string(hashed),
	}

	if err := models.SaveUser(&user); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	// log user in after account creation
	if err := utils.Authenticate(w, req.Username, req.Password); err != nil {
		http.Error(w, "Failed to log in after account creation", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Account created and logged in successfully"))
}
