package usercontext

import "context"

// contextKey is a custom type to avoid collisions in context keys.
type contextKey string

// Key is the context key used to store user ID.
const Key contextKey = "userid"

// GetUserID retrieves the user ID from the given context.
func GetUserID(ctx context.Context) (userID string, ok bool) {
	userID, ok = ctx.Value(Key).(string)
	return
}

// WithID returns a new context with the given user ID.
func WithID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, Key, userID)
}
