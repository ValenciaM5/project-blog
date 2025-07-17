package main

import (
	"log"
	"net/http"
	"os"

	"project-blog/db"
	"project-blog/handlers"
	"project-blog/repository"
	"project-blog/routes"
	"project-blog/services"

	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file, using default environment variables")
	}

	// Initialize Database
	db.InitDB()

	// Initialize Repositories
	userRepo := repository.NewUserRepository(db.DB)
	postRepo := repository.NewPostRepository(db.DB)

	// Initialize Services
	userService := services.NewUserService(userRepo)
	postService := services.NewPostService(postRepo)

	// Initialize Handlers
	userHandler := handlers.NewUserHandler(userService)
	postHandler := handlers.NewPostHandler(postService)

	// Setup Router
	router := routes.SetupRouter(userHandler, postHandler)

	// Get port from environment variable or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start Server
	log.Printf("Server running at http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
