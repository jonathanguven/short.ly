package models

import (
	"shortly/internal/database"
)

// User struct for the database model
type User struct {
	ID           uint   `gorm:"primaryKey"`
	Username     string `gorm:"unique"`
	PasswordHash string
}

// SaveUser saves a new User to the database
func SaveUser(user *User) error {
	return database.DB.Create(user).Error
}

// FindUserByUsername finds a User by its username
func FindUserByUsername(username string) (*User, error) {
	var user User
	result := database.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
