package main

import (
	"fmt"
	"log"
	"net/http"
	"shortly/internal/database"
	"shortly/internal/handlers"

	"shortly/internal/middlewares"
	"shortly/internal/models"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Go URL shortener")
}

func main() {
	database.InitializeDB()
	database.DB.AutoMigrate(&models.URL{}, &models.User{})

	http.HandleFunc("/", greet)
	http.Handle("/shorten", middlewares.Authenticate(http.HandlerFunc(handlers.HandleShorten)))
	http.HandleFunc("/s/", handlers.HandleRedirect)
	http.HandleFunc("/login", handlers.HandleLogin)
	http.HandleFunc("/create-account", handlers.HandleCreateUser)
	http.HandleFunc("/list", handlers.HandleListURLs)

	log.Println("Starting server on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
