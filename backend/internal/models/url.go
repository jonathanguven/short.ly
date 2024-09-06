package models

import (
	"shortly/internal/database"
	"time"
)

// URL struct for the database model
type URL struct {
	ID        uint   `gorm:"primaryKey"`
	Alias     string `gorm:"unique"`
	URL       string
	CreatedAt time.Time
	ExpiresAt *time.Time
	UserID    uint
}

// SaveURL saves a new URL to the database
func SaveURL(url *URL) error {
	return database.DB.Create(url).Error
}

// FindURL finds a URL by its alias
func FindURL(alias string) (*URL, error) {
	var url URL
	result := database.DB.Where("alias = ?", alias).First(&url)
	if result.Error != nil {
		return nil, result.Error
	}
	return &url, nil
}
