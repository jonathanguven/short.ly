package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func InitLogger() {
	log.SetLevel(log.InfoLevel)
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.JSONFormatter{})
	log.Info("Logger initialized")
}
