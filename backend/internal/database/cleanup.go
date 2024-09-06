package database

import (
	"log"
	"shortly/internal/models"
	"time"

	"gorm.io/gorm"
)

// background job to periodically remove expired URLs
func StartCleanup(db *gorm.DB) {
	ticker := time.NewTicker(24 * time.Hour)
	go func() {
		for range ticker.C {
			cleanupExpiredURLs(db)
		}
	}()
}

// remove URLs older than expiration date
func cleanupExpiredURLs(db *gorm.DB) {
	expired := db.Where("user_id AND expires_at < ?", time.Now()).Delete(&models.URL{})
	if expired.Error != nil {
		log.Printf("Error cleaning up expired URLs: %v", expired.Error)
		return
	}
	if expired.RowsAffected > 0 {
		log.Printf("Cleaned up %d expired URLs", expired.RowsAffected)
	}
}
