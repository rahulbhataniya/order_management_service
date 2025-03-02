package main

import (
	"fmt"
	"log"
	"order-management/models"
	"order-management/config" // Updated import
)

func main() {
	// Connect to the database
	db := config.ConnectDB()

	// Run AutoMigrate to create tables automatically
	err := db.AutoMigrate(&models.Order{})
	if err != nil {
		log.Fatalf("Error during migration: %v", err)
	}

	fmt.Println("✅ Migration completed successfully!")
}
