package authentication

import (
	"context"
	"firebase.google.com/go/v4/auth"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"mai.today/api/gen/authentication"
	"mai.today/user"
)

type Authentication struct {
	firebaseAuth *auth.Client
	mongo        *mongo.Client
}

func NewAuthentication(client *auth.Client, mongoClient *mongo.Client) *Authentication {
	return &Authentication{
		firebaseAuth: client,
		mongo:        mongoClient,
	}
}

func (a *Authentication) SignIn(c context.Context, payload *authentication.SignInPayload) (err error) {

	token, err := a.firebaseAuth.VerifyIDToken(c, payload.FirebaseIDToken)
	if err != nil {
		log.Printf("error verifying ID token: %v\n", err)
		return authentication.MakeTokenError(err)
	}
	log.Printf("Verified ID token: %v\n", token)

	var user user.IUser = user.NewUser(a.mongo)
	userInfo := user.FindUserByFireBaseUID(token.UID)

	if userInfo == nil {
		log.Printf("User not found: %v\n", "creating user...")
		err := user.CreateUser(token.UID)
		if err != nil {
			log.Printf("error creating user: %v\n", err)
			//TODO change to create user error
			return authentication.MakeTokenError(err)
		}
	}

	return nil
}
