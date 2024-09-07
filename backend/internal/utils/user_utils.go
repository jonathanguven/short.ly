package utils

import (
	"shortly/internal/database"
	"shortly/internal/models"
)

// SaveUser saves a new User to the database
func SaveUser(user *models.User) error {
	return database.DB.Create(user).Error
}

// FindUserByUsername finds a User by its username
func FindUserByUsername(username string) (*models.User, error) {
	var user models.User
	result := database.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func FindURLsByUser(userID uint) ([]models.URL, error) {
	var urls []models.URL
	result := database.DB.Where("user_id = ?", userID).Find(&urls)
	return urls, result.Error
}
