package models

import (
	"shortly/internal/database"
	"time"
)

// URL struct for the database model
type URL struct {
	ID          uint   `gorm:"primaryKey"`
	ShortID     string `gorm:"unique"`
	OriginalURL string
	CreatedAt   time.Time
	ExpiresAt   *time.Time
	UserID      uint
}

// SaveURL saves a new URL to the database
func SaveURL(url *URL) error {
	return database.DB.Create(url).Error
}

// FindURLByShortID finds a URL by its short ID
func FindURLByShortID(shortID string) (*URL, error) {
	var url URL
	result := database.DB.Where("short_id = ?", shortID).First(&url)
	if result.Error != nil {
		return nil, result.Error
	}
	return &url, nil
}
