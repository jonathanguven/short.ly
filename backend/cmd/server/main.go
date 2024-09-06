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
	database.InitializeDB()
	database.DB.AutoMigrate(&models.URL{}, &models.User{})

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Go URL shortener")
	})
	r.Handle("/shorten", middlewares.Authenticate(http.HandlerFunc(handlers.HandleShorten)))
	r.HandleFunc("/s/{alias}", handlers.HandleRedirect).Methods("GET")
	r.HandleFunc("/login", handlers.HandleLogin)
	r.HandleFunc("/create-account", handlers.HandleCreateUser)
	r.HandleFunc("/urls", handlers.HandleListURLs)

	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
