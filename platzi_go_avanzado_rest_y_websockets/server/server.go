package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"platzi/go/rest_websockets/database"
	"platzi/go/rest_websockets/repository"

	"github.com/gorilla/mux"
)

type Config struct {
	Port        string
	JwtSecret   string
	DatabaseURL string
}

type Server interface {
	Config() *Config
}

type Broker struct {
	config *Config
	router *mux.Router
}

func (b *Broker) Config() *Config {
	return b.config
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {

	if config.JwtSecret == "" {
		return nil, errors.New("JwtSecret is required")
	}

	if config.DatabaseURL == "" {
		return nil, errors.New("DatabaseURL is required")
	}

	if config.Port == "" {
		return nil, errors.New("port is required")
	}

	return &Broker{
		config: config,
		router: mux.NewRouter(),
	}, nil
}

func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	binder(b, b.router)

	repo, err := database.NewPostgresRepository(b.config.DatabaseURL)
	if err != nil {
		log.Fatalf("Error creating repository: %v", err)
	}

	repository.SetRepository(repo)

	log.Println("Server is running on port", b.config.Port)
	http.Handle("/", b.router)
	if err := http.ListenAndServe(b.config.Port, b.router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

// func main() {
// 	config := &Config{
// 		Port:        ":8080",
// 		JwtSecret:   "secret",
// 		DatabaseURL: "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable",
// 	}
// 	server, err := NewServer(context.Background(), config)
// 	if err != nil {
// 		log.Fatalf("Error creating server: %v", err)
// 	}
// 	server.Start(func(s Server, r *mux.Router) {
// 		r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 			w.Write([]byte("welcome"))
// 		}).Methods("GET")
// 	})
// }
