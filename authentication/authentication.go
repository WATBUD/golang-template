package authentication

import (
	"context"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"goa.design/goa/v3/security"

	"mai.today/api/gen/authentication"
	"mai.today/authentication/internal"
	"mai.today/database/mongodb"
	usercontext "mai.today/foundation/context/user"
	"mai.today/user"
)

var (
	// once ensures instance initialization is performed exactly once.
	once sync.Once

	// instance holds the singleton Authentication instance.
	instance Authentication
)

// Instance returns a singleton instance of Authentication.
func Instance() Authentication {
	once.Do(func() {
		instance = newAuthentication()
	})
	return instance
}

type Authentication struct {
	verifier internal.TokenVerifier
	mongo    *mongo.Client
}

func newAuthentication() Authentication {
	mongoClient := mongodb.Instance()
	firebase := InitFirebase("../maitoday-168-dev-fireBase.json")

	return Authentication{
		firebase,
		mongoClient,
	}
}

func (a Authentication) SignIn(c context.Context, payload *authentication.SignInPayload) (err error) {
	token, err := a.verifier.VerifyIDToken(c, payload.FirebaseIDToken)
	if err != nil {
		return authentication.MakeTokenError(err)
	}

	var user user.IUser = user.NewUser(a.mongo)
	userInfo := user.FindUserByFireBaseUID(token.UID)

	if userInfo == nil {
		err := user.CreateUser(token.UID)
		if err != nil {
			return err
		}
	}

	return nil
}

// VerifyToken verifies the provided JWT token and returns the extracted user ID or an error
//
// ctx: context.Context containing the incoming request context
// token: string representing the JWT token to be verified
//
// Returns:
//
//	string: The user ID extracted from the verified token, empty string on error
//	error: Any error encountered during token verification
func (a Authentication) VerifyToken(ctx context.Context, token string) (userID string, err error) {
	t, err := a.verifier.VerifyIDToken(ctx, token)
	if err != nil {
		return "", err
	}

	return t.UID, err
}

// JWTAuth verifies the JWT token and injects user ID into context
//
// ctx: context.Context containing the incoming request context
// token: string representing the JWT token to be verified (optional, depending on implementation)
// schema: *security.JWTScheme (used for schema validation, assumed to be provided)
//
// Returns:
//
//	context.Context: The context with the user ID added if successful, original context otherwise
//	error: Any error encountered during token verification or context manipulation
func (a Authentication) JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error) {
	id, err := a.VerifyToken(ctx, token)
	if err != nil {
		return ctx, authentication.MakeInvalidToken(err)
	}

	return usercontext.WithID(ctx, id), nil
}
