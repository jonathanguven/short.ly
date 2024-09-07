package main

import (
	"net/http"
	"shortly/internal/database"
	"shortly/internal/handlers"

	"shortly/internal/middlewares"
	"shortly/internal/models"

	"shortly/internal/metrics"
	"shortly/internal/utils"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

	r := mux.NewRouter()

	r.Handle("/shorten", middlewares.Authenticate(http.HandlerFunc(handlers.HandleShorten))).Methods("POST")
	r.HandleFunc("/s/{alias}", handlers.HandleRedirect).Methods("GET")
	r.HandleFunc("/login", handlers.HandleLogin).Methods("POST")
	r.HandleFunc("/create-account", handlers.HandleCreateUser).Methods("POST")
	r.HandleFunc("/urls", handlers.HandleListURLs).Methods("GET")
	r.Handle("/metrics", promhttp.Handler())

	log.Info("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
