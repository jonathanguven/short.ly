package utils

import (
	"errors"
	"shortly/internal/models"
)

// verify logged in user: user can only modify/delete their own urls
func VerifyUser(userID uint, url *models.URL) error {
	if url.UserID != userID {
		return errors.New("unauthorized: user does not own this URL")
	}

	return nil
}
