package utils

import (
	"net/http"
	"shortly/internal/models"

	"golang.org/x/crypto/bcrypt"
)

// validates user name and password and logs user in
func Authenticate(w http.ResponseWriter, username, password string) error {
	// get user from db using username
	user, err := models.FindUserByUsername(username)
	if err != nil {
		return err
	}

	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return err
	}

	// generate JWT token
	token, err := GenerateJWT(user.ID)
	if err != nil {
		return err
	}

	// set JWT token in HTTP-only cookie
	SetCookie(w, token)

	return nil
}
