package handlers

import (
	"encoding/json"
	"net/http"
	"reminder_backend/services"
	"reminder_backend/utils"

	"github.com/golang-jwt/jwt"
	"google.golang.org/api/calendar/v3"
)

func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	tokenStr := r.Header.Get("Authorization")
	claims := &utils.Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return utils.JwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var event *calendar.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	event.Start = &calendar.EventDateTime{DateTime: "2024-06-19T10:00:00-07:00", TimeZone: "America/Los_Angeles"}
	event.End = &calendar.EventDateTime{DateTime: "2024-06-19T10:25:00-07:00", TimeZone: "America/Los_Angeles"}

	calendarID := "primary"
	event, err = services.CalendarService.Events.Insert(calendarID, event).Do()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Schedule notification (Firebase Cloud Messaging) 5 minutes before the event
	services.SendPushNotification(event.Start.DateTime, "Event Reminder", "You have an event starting in 5 minutes.")

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(event)
}
