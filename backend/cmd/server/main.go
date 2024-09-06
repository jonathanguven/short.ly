package main

import (
	"fmt"
	"log"
	"net/http"
	"shortly/internal/database"
	"shortly/internal/handlers"

	"shortly/internal/middlewares"
	"shortly/internal/models"

	"github.com/gorilla/mux"
)

func main() {
	// connect to database
	database.InitializeDB()
	database.DB.AutoMigrate(&models.URL{}, &models.User{})

	// initialize cleanup cron job
	database.StartCleanup(database.DB)

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Go URL shortener")
	}).Methods("GET")
	r.Handle("/shorten", middlewares.Authenticate(http.HandlerFunc(handlers.HandleShorten))).Methods("POST")
	r.HandleFunc("/s/{alias}", handlers.HandleRedirect).Methods("GET")
	r.HandleFunc("/login", handlers.HandleLogin).Methods("POST")
	r.HandleFunc("/create-account", handlers.HandleCreateUser).Methods("POST")
	r.HandleFunc("/urls", handlers.HandleListURLs).Methods("GET")

	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
