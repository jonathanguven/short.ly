package utils

import (
	"shortly/internal/database"
	"shortly/internal/models"
)

// SaveURL saves a new URL to the database
func SaveURL(url *models.URL) error {
	return database.DB.Save(url).Error
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

// only update the click count column
func UpdateClickCount(id uint, clickCount int) error {
	return database.DB.Model(&models.URL{}).Where("id = ?", id).Update("click_count", clickCount).Error
}
