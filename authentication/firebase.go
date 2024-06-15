package authentication

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
	"log"
)

func InitFirebase(credentialsFilePath string) *auth.Client {
	var firebaseAuth *auth.Client
	opt := option.WithCredentialsFile(credentialsFilePath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	// Access auth service from the default app
	firebaseAuth, err = app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	log.Printf("Firebase Auth initialized\n")
	return firebaseAuth
}
