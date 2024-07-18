package authentication

import (
	"context"
	"errors"
	"testing"

	"firebase.google.com/go/v4/auth"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"mai.today/authentication/internal/mock"
	usercontext "mai.today/foundation/context/user"
)

func TestAuthentication_VerifyToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockVerifier := mock.NewMockTokenVerifier(ctrl)
	mockMongo := &mongo.Client{}

	ctx := context.Background()
	token := "someToken"
	expected := &auth.Token{
		UID: "user123",
	}

	t.Run("successful token verification", func(t *testing.T) {
		mockVerifier.EXPECT().VerifyIDToken(ctx, token).Return(expected, nil)

		a := &Authentication{mockVerifier, mockMongo}
		result, err := a.VerifyToken(ctx, token)

		assert.NoError(t, err)
		assert.Equal(t, expected.UID, result)
	})

	t.Run("failed token verification", func(t *testing.T) {
		mockVerifier.EXPECT().VerifyIDToken(ctx, token).Return(expected, errors.New("invalid token"))

		a := &Authentication{mockVerifier, mockMongo}
		_, err := a.VerifyToken(ctx, token)

		assert.Error(t, err)
	})

}

func TestAuthentication_JWTAuth(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockVerifier := mock.NewMockTokenVerifier(ctrl)
	mockMongo := &mongo.Client{}

	ctx := context.Background()
	token := "someToken"
	expected := &auth.Token{
		UID: "user123",
	}

	t.Run("successful token verification", func(t *testing.T) {
		mockVerifier.EXPECT().VerifyIDToken(ctx, token).Return(expected, nil)

		a := &Authentication{mockVerifier, mockMongo}
		result, err := a.JWTAuth(ctx, token, nil)

		id, ok := usercontext.GetUserID(result)
		assert.True(t, ok)
		assert.Equal(t, expected.UID, id)
		assert.NoError(t, err)
	})

	t.Run("failed token verification", func(t *testing.T) {
		mockVerifier.EXPECT().VerifyIDToken(ctx, token).Return(expected, errors.New("invalid token"))

		a := &Authentication{mockVerifier, mockMongo}
		result, err := a.JWTAuth(ctx, token, nil)

		id, ok := usercontext.GetUserID(result)
		assert.False(t, ok)
		assert.Empty(t, id)
		assert.Error(t, err)
	})
}
