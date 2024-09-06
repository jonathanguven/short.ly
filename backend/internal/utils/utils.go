package utils

import (
	"math/rand"
	"shortly/internal/database"
	"shortly/internal/models"
	"time"
)

// returns all shortened URLs from the database
func FindAllURLs() ([]models.URL, error) {
	var urls []models.URL
	result := database.DB.Find(&urls)
	return urls, result.Error
}

// generate a random string of length 5
func GenerateHash() string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	// random generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]rune, 5)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}
