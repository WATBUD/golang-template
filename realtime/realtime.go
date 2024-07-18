package realtime

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/centrifugal/gocent/v3"
	"mai.today/realtime/internal/centrifugo"
)

var (
	// once ensures instance initialization is performed exactly once.
	once sync.Once

	// instance holds the singleton Realtime instance.
	instance Realtime
)

// Realtime represents a Centrifugo client instance.
type Realtime struct {
	centrifugo centrifugo.Client
}

// Instance returns a singleton instance of Realtime.
func Instance() Realtime {
	once.Do(func() {
		instance = newRealtime(defaultOrEnv())
	})
	return instance
}

// Broadcast sends data to specified channels using Centrifugo.
func (r *Realtime) Broadcast(ctx context.Context, channels []string, data interface{}) (gocent.BroadcastResult, error) {
	d, err := encodeToBytes(data)
	if err != nil {
		return gocent.BroadcastResult{}, err
	}

	return r.centrifugo.Broadcast(ctx, channels, d)
}

// BroadcastToUsers sends data to specified user channels using Centrifugo.
func (r *Realtime) BroadcastToUsers(ctx context.Context, userIDs []string, data interface{}) (gocent.BroadcastResult, error) {
	c := make([]string, len(userIDs))
	for i, userID := range userIDs {
		c[i] = addUserPrefix(userID)
	}

	return r.Broadcast(ctx, c, data)
}

// Publish publishes data to a specific channel using Centrifugo.
func (r *Realtime) Publish(ctx context.Context, channel string, data interface{}) (gocent.PublishResult, error) {
	d, err := encodeToBytes(data)
	if err != nil {
		return gocent.PublishResult{}, err
	}

	return r.centrifugo.Publish(ctx, channel, d)
}

// PublishToUser publishes data to a specific user channel using Centrifugo.
func (r *Realtime) PublishToUser(ctx context.Context, userID string, data interface{}) (gocent.PublishResult, error) {
	c := addUserPrefix(userID)
	return r.Publish(ctx, c, data)
}

// addUserPrefix prefixes user IDs with "#" as expected by Centrifugo.
func addUserPrefix(userID string) string {
	return "#" + userID
}

// encodeToBytes encodes data to JSON bytes.
func encodeToBytes(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

// newRealtime creates a new Realtime instance with given Centrifugo address and key.
func newRealtime(addr, key string) Realtime {
	c := gocent.New(gocent.Config{Addr: addr, Key: key})
	return Realtime{centrifugo: c}
}
