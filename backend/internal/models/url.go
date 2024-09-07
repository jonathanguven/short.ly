package models

import (
	"time"
)

// URL struct for the database model
type URL struct {
	ID         uint   `gorm:"primaryKey"`
	Alias      string `gorm:"unique"`
	URL        string
	CreatedAt  time.Time
	ExpiresAt  *time.Time
	UserID     uint
	ClickCount int `gorm:"default:0"`
}
