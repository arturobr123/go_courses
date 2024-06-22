package main

import (
	"log"
	"net/http"
	"reminder_backend/handlers"
	"reminder_backend/services"

	"github.com/gorilla/mux"
)

func main() {
	// services.InitFirebase()
	services.InitGoogleCalendar()

	r := mux.NewRouter()
	r.HandleFunc("/signup", handlers.SignupHandler).Methods("POST")
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	r.HandleFunc("/create-event", handlers.CreateEventHandler).Methods("POST")

	log.Println("Server is ready to serve")
	log.Fatal(http.ListenAndServe(":8080", r))
}
