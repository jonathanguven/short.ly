package utils

import (
	"math/rand"
	"time"
)

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
