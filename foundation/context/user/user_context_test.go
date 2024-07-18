package usercontext

import (
	"context"
	"testing"
)

func TestGetUserID(t *testing.T) {
	tests := []struct {
		name     string
		ctx      context.Context
		expected string
		ok       bool
	}{
		{
			name:     "UserID present",
			ctx:      WithID(context.Background(), "12345"),
			expected: "12345",
			ok:       true,
		},
		{
			name:     "UserID not present",
			ctx:      context.Background(),
			expected: "",
			ok:       false,
		},
		{
			name:     "UserID of wrong type",
			ctx:      context.WithValue(context.Background(), Key, 12345),
			expected: "",
			ok:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userID, ok := GetUserID(tt.ctx)
			if userID != tt.expected || ok != tt.ok {
				t.Errorf("UserIDFromContext() = %v, %v; want %v, %v", userID, ok, tt.expected, tt.ok)
			}
		})
	}
}

func TestWithID(t *testing.T) {
	str12345 := "12345"
	ctx := context.Background()

	newCtx := WithID(ctx, str12345)
	userID, ok := GetUserID(newCtx)

	if !ok || userID != str12345 {
		t.Errorf("WithValue() or FromContext() failed: got %v, %v; want %v, true", userID, ok, str12345)
	}
}
