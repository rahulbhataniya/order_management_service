package config

import (
	"fmt"
	"log"
	"os"
	"order-management/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB establishes a connection to MySQL using GORM
func ConnectDB() *gorm.DB {
	// Fetch environment variables directly
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Construct the DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	// Connect to the database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Database connection failed: %v", err)
	}

	fmt.Println("✅ Database connected successfully")

	// Auto-Migrate models
	err = db.AutoMigrate(&models.Order{})
	if err != nil {
		log.Fatalf("❌ AutoMigrate failed: %v", err)
	}

	DB = db
	return db
}
