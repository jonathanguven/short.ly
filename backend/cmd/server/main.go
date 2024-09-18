package main

import (
	"net/http"
	"os"
	"shortly/internal/database"
	"shortly/internal/handlers"

	"shortly/internal/middlewares"
	"shortly/internal/models"

	"shortly/internal/metrics"
	"shortly/internal/utils"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

func main() {
	// set up logger
	utils.InitLogger()

	// set up prometheus client
	metrics.Init()

	// connect to database
	database.InitializeDB()
	database.DB.AutoMigrate(&models.URL{}, &models.User{})

	// initialize cleanup cron job
	database.StartCleanup(database.DB)

	domain := os.Getenv("DOMAIN")

	r := mux.NewRouter()

	r.Handle("/shorten", middlewares.Authenticate(http.HandlerFunc(handlers.HandleShorten))).Methods("POST")
	r.Handle("/urls/{alias}", middlewares.Authenticate(http.HandlerFunc(handlers.HandleUpdateURL))).Methods("PUT")
	r.Handle("/urls/{alias}", middlewares.Authenticate(http.HandlerFunc(handlers.HandleDeleteURL))).Methods("DELETE")
	r.HandleFunc("/s/{alias}", handlers.HandleRedirect).Methods("GET")
	r.HandleFunc("/login", handlers.HandleLogin).Methods("POST")
	r.HandleFunc("/logout", handlers.HandleLogout).Methods("POST")
	r.HandleFunc("/create-account", handlers.HandleCreateUser).Methods("POST")
	r.Handle("/urls", middlewares.Authenticate(http.HandlerFunc(handlers.HandleListURLs))).Methods("GET")
	r.Handle("/metrics", promhttp.Handler())

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{domain},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}).Handler(r)

	log.Info("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}
