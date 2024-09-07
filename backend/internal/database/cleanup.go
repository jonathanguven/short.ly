package database

import (
	"shortly/internal/models"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// background job to periodically remove expired URLs
func StartCleanup(db *gorm.DB) {
	ticker := time.NewTicker(24 * time.Hour)
	go func() {
		for range ticker.C {
			log.Println("Starting cleanup task...")
			cleanupExpiredURLs(db)
		}
	}()
}

// remove URLs older than expiration date
func cleanupExpiredURLs(db *gorm.DB) {
	expired := db.Where("user_id = 0 AND expires_at < ?", time.Now().UTC()).Delete(&models.URL{})

	if expired.Error != nil {
		log.Printf("Error cleaning up expired URLs: %v", expired.Error)
		return
	}

	if expired.RowsAffected > 0 {
		log.Printf("Cleaned up %d expired URLs\n", expired.RowsAffected)
	} else {
		log.Printf("No expired URLs found for cleanup\n")
	}
	log.Println("Cleanup task completed")
}
