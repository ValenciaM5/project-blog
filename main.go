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

	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Error loading .env file, using default environment variables")
	}

	db.InitDB()

	userRepo := repository.NewUserRepository(db.DB)
	postRepo := repository.NewPostRepository(db.DB)

	userService := services.NewUserService(userRepo)
	postService := services.NewPostService(postRepo)

	userHandler := handlers.NewUserHandler(userService)
	postHandler := handlers.NewPostHandler(postService)

	router := routes.SetupRouter(userHandler, postHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running at http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

// Live App:https://project-blog-oxkc.onrender.com
