package utils

import (
	"shortly/internal/database"
	"shortly/internal/models"
)

// SaveURL saves a new URL to the database
func SaveURL(url *models.URL) error {
	return database.DB.Create(url).Error
}

// FindURL finds a URL by its alias
func FindURL(alias string) (*models.URL, error) {
	var url models.URL
	result := database.DB.Where("alias = ?", alias).First(&url)
	if result.Error != nil {
		return nil, result.Error
	}
	return &url, nil
}
