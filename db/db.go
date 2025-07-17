package db

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"project-blog/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file, using default environment variables")
	}

	// Get database credentials from environment
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPortStr := os.Getenv("DB_PORT")
	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		log.Fatalf("Invalid DB_PORT: %v. Ensure DB_PORT is set to 5433 in .env", err)
	}
	dbName := os.Getenv("DB_NAME")

	// Validate required environment variables
	if dbHost == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		log.Fatal("Missing required environment variables: DB_HOST, DB_USER, DB_PASSWORD, or DB_NAME")
	}

	// First, connect to the postgres database to create blog_db if needed
	adminDSN := fmt.Sprintf("host=%s user=%s password=%s dbname=postgres port=%d sslmode=disable",
		dbHost, dbUser, dbPassword, dbPort)
	adminDB, err := gorm.Open(postgres.Open(adminDSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to postgres database:", err)
	}

	// Create blog_db if it doesn't exist
	sqlDB, err := adminDB.DB()
	if err != nil {
		log.Fatal("Failed to get database handle:", err)
	}

	// Connect to the target database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Enable uuid-ossp extension in the target database
	sqlDB, err = db.DB()
	if err != nil {
		log.Fatal("Failed to get database handle for blog_db:", err)
	}
	_, err = sqlDB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	if err != nil {
		log.Fatal("Failed to enable uuid-ossp extension in blog_db:", err)
	}

	// Auto-migrate models
	err = db.AutoMigrate(&models.User{}, &models.Post{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	DB = db
	log.Println("Database connected successfully")
}
