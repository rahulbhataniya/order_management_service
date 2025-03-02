package utils

import (
	"log"
	"os"
)

// Logger instance
var Logger *log.Logger

func init() {
	// Create a log file
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	// Initialize the logger
	Logger = log.New(file, "ORDER-MANAGEMENT: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// LogInfo logs informational messages
func LogInfo(message string) {
	Logger.Println("[INFO]: " + message)
}

// LogError logs error messages
func LogError(err error) {
	if err != nil {
		Logger.Println("[ERROR]: " + err.Error())
	}
}
