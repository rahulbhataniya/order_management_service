package config

import (
	"fmt"
	"log"
	"os"
	"order-management/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

// LoadEnv loads environment variables from .env file
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  No .env file found, using system environment variables")
	}
}

// ConnectDB establishes a connection to MySQL using GORM
func ConnectDB() *gorm.DB {
	LoadEnv() // Load environment variables

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Database connection failed: %v", err)
	}

	fmt.Println("✅ Database connected successfully")

	// **Auto-Migrate models**
	err = db.AutoMigrate(&models.Order{})
	if err != nil {
		log.Fatalf("❌ AutoMigrate failed: %v", err)
	}

	DB = db
	return db
}
