package services

import (
	"context"
	"log"
	"time"

	"firebase.google.com/go/messaging"
)

func SendPushNotification(eventTime, title, body string) {
	// Calculate the time to send the notification (5 minutes before event)
	eventTimeParsed, _ := time.Parse(time.RFC3339, eventTime)
	notificationTime := eventTimeParsed.Add(-5 * time.Minute)

	// Set up a goroutine to send the notification at the calculated time
	go func() {
		time.Sleep(time.Until(notificationTime))
		message := &messaging.Message{
			Notification: &messaging.Notification{
				Title: title,
				Body:  body,
			},
			Topic: "event_reminders",
		}

		_, err := MessagingClient.Send(context.Background(), message)
		if err != nil {
			log.Printf("error sending push notification: %v\n", err)
		}
	}()
}
