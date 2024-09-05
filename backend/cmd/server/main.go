package main

import (
	"log"
	"net/http"
	"shortly/internal/database"
	"shortly/internal/models"
)

func main() {
	database.InitializeDB()
	database.DB.AutoMigrate(&models.URL{}, &models.User{})

	// http.HandleFunc("/shorten", handleShorten)
	// http.HandleFunc("/redirect/", handleRedirect)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
