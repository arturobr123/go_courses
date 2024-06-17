package main

import (
	"context"
	"log"
	"os"
	"platzi/go/rest_websockets/handlers"
	"platzi/go/rest_websockets/middleware"
	"platzi/go/rest_websockets/server"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	port := os.Getenv("PORT")
	jwtSecret := os.Getenv("JWT_SECRET")
	databaseURL := os.Getenv("DATABASE_URL")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port:        port,
		JwtSecret:   jwtSecret,
		DatabaseURL: databaseURL,
	})

	if err != nil {
		log.Fatalf("Error creating server: %v", err)
	}

	s.Start(BindRoutes)
}

func BindRoutes(s server.Server, r *mux.Router) {
	// authentication if need it
	r.Use(middleware.CheckAuthMiddleware(s))

	r.HandleFunc("/", handlers.HomeHandler(&s)).Methods("GET")
	r.HandleFunc("/signup", handlers.SignUpHandler(s)).Methods("POST")
	r.HandleFunc("/login", handlers.LogInHandler(s)).Methods("POST")
	r.HandleFunc("/me", handlers.MeHandler(s)).Methods("GET")
	r.HandleFunc("/posts", handlers.InsertPostHandler(s)).Methods("POST")
	r.HandleFunc("/posts/{id}", handlers.GetPostByIdHandler(s)).Methods("GET")
	r.HandleFunc("/posts/{id}", handlers.UpdatePostHandler(s)).Methods("PUT")
	r.HandleFunc("/posts/{id}", handlers.DeletePostHandler(s)).Methods("DELETE")
	r.HandleFunc("/posts", handlers.ListPostsHandler(s)).Methods("GET") // localhost:5050/posts?page=1
}

// Capilla Jesus buen pastor
