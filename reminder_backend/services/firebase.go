package services

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

var (
	FirebaseApp     *firebase.App
	FirebaseAuth    *auth.Client
	FirestoreClient *firestore.Client
	MessagingClient *messaging.Client
)

func InitFirebase() {
	opt := option.WithCredentialsFile("serviceAccountKey.json")

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing firebase app: %v\n", err)
	}

	authClient, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error initializing firebase auth: %v\n", err)
	}

	firestore, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalf("error initializing firestore: %v\n", err)
	}

	messagingClient, err := app.Messaging(context.Background())
	if err != nil {
		log.Fatalf("error initializing messaging client: %v\n", err)
	}

	FirebaseApp = app
	FirebaseAuth = authClient
	FirestoreClient = firestore
	MessagingClient = messagingClient
}
