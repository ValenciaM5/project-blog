Go Blog Platform
A RESTful blog platform built with Golang, GORM, PostgreSQL, and JWT authentication, deployed on Render.
Technologies Used

Go 1.21
GORM (PostgreSQL driver)
Gorilla Mux
JWT for authentication
bcrypt for password hashing
godotenv for environment variables
Render for deployment 

The project is deployed on Render:
[Click here to view the app] https://project-blog-oxkc.onrender.com


Setup Instructions
Prerequisites

Go 1.21+
PostgreSQL
Render account for deployment

Local Setup

Clone the repository:

git clone <repository-url>
cd go-blog


Copy .env.example to .env and fill in your configuration:

cp .env.example .env


Install dependencies:

go mod download


Set up PostgreSQL database and ensure it's running.

Run the application:


go run main.go

The server will start at http://localhost:8080 (or the port specified in .env).
API Endpoints

POST /register - Register a new user
POST /login - Login and get JWT token
GET /api/me - Get authenticated user info
POST /api/posts - Create a new post
GET /api/posts - List all posts
GET /api/posts/:id - Get a specific post
PUT /api/posts/:id - Update a post
DELETE /api/posts/:id - Delete a post

Deployment Instructions (Render)

Create a new Web Service on Render.
Connect your GitHub repository.
Configure the following in Render:
Build Command: go build -o blog
Start Command: ./blog
Environment Variables:
DB_HOST: Your PostgreSQL host
DB_USER: Your PostgreSQL user
DB_PASSWORD: Your PostgreSQL password
DB_NAME: Your database name
DB_PORT: Your PostgreSQL port
JWT_SECRET: A secure random string
PORT: Application port (usually 8080)




Deploy the service.

Live App Link
[To be added after deployment]
Notes

Ensure your PostgreSQL database is accessible from Render
Use a strong JWT_SECRET for production
The API requires Bearer token authentication for protected routes
Only post authors can update/delete their own posts
