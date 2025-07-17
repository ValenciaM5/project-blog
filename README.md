This is a RESTful blog platform built using Go, GORM, PostgreSQL, and JWT authentication. It is deployed on Render.

Technologies Used
Go 1.21

GORM (PostgreSQL driver)

Gorilla Mux

JWT (for authentication)

bcrypt (for password hashing)

godotenv (for environment variable management)

Render (for deployment)

Live App:
https://project-blog-oxkc.onrender.com

Setup Instructions
Prerequisites
Go 1.21 or later

PostgreSQL

A Render account

Local Setup
Clone the repository:


git clone <repository-url>
cd go-blog
Copy the example environment file and update the values:

cp .env.example .env
Download dependencies:


go mod download
Set up your PostgreSQL database and make sure it is running.

Run the application:
go run main.go

The server will start at http://localhost:8080 or the port specified in the .env file.

API Endpoints
Authentication
POST /register – Register a new user

POST /login – Log in and receive a JWT token

Users
GET /api/me – Get information about the logged-in user

Posts
POST /api/posts – Create a new blog post

GET /api/posts – Get all blog posts

GET /api/posts/:id – Get a single post by ID

PUT /api/posts/:id – Update a post (only if logged in as the user)

DELETE /api/posts/:id – Delete a post (only if logged in as the user)

Deployment on Render
Create a new Web Service on Render.

Connect your GitHub repository.

Use the default settings.


Environment Variables:


DB_HOST=your-database-host
DB_USER=your-database-username
DB_PASSWORD=your-database-password
DB_NAME=your-database-name
DB_PORT=your-database-port
JWT_SECRET=your-jwt-secret
PORT=8080
Deploy the service.

Notes
Make sure your PostgreSQL database allows external access so Render can connect to it.

Use a secure value for JWT_SECRET in production.

Protected routes require a Bearer token in the Authorization header.

Only the user who created a post can update or delete it.

