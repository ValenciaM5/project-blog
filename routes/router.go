package routes

import (
	"context"
	"net/http"
	"os"
	"strings"

	"project-blog/handlers"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func SetupRouter(userHandler handlers.UserHandler, postHandler handlers.PostHandler) *mux.Router {
	router := mux.NewRouter()

	// Auth routes
	router.HandleFunc("/register", userHandler.Register).Methods("POST")
	router.HandleFunc("/login", userHandler.Login).Methods("POST")

	// Protected routes
	api := router.PathPrefix("/api").Subrouter()
	api.Use(AuthMiddleware)
	api.HandleFunc("/me", userHandler.Me).Methods("GET")
	api.HandleFunc("/posts", postHandler.Create).Methods("POST")
	api.HandleFunc("/posts", postHandler.FindAll).Methods("GET")
	api.HandleFunc("/posts/{id}", postHandler.FindByID).Methods("GET")
	api.HandleFunc("/posts/{id}", postHandler.Update).Methods("PUT")
	api.HandleFunc("/posts/{id}", postHandler.Delete).Methods("DELETE")

	return router
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrNotSupported
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", claims["user_id"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
