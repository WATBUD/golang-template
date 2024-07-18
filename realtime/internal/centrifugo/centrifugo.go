package centrifugo

import (
	context "context"
	"net/http"

	"github.com/centrifugal/gocent/v3"
)

// mockgen -source=internal/centrifugo/centrifugo.go -destination=internal/centrifugo/mock_centrifugo.go -package=centrifugo

type Client interface {
	// SetHTTPClient allows to set custom http Client to use for requests. Not goroutine-safe.
	SetHTTPClient(httpClient *http.Client)
	// Pipe allows to create new Pipe to send several commands in one HTTP request.
	Pipe() *gocent.Pipe
	// Publish allows to publish data to channel.
	Publish(ctx context.Context, channel string, data []byte, opts ...gocent.PublishOption) (gocent.PublishResult, error)
	// Broadcast allows to broadcast the same data into many channels..
	Broadcast(ctx context.Context, channels []string, data []byte, opts ...gocent.PublishOption) (gocent.BroadcastResult, error)
	// Subscribe allow subscribing user to a channel (using server-side subscriptions).
	Subscribe(ctx context.Context, channel, user string, opts ...gocent.SubscribeOption) error
	// Unsubscribe allows to unsubscribe user from channel.
	Unsubscribe(ctx context.Context, channel, user string, opts ...gocent.UnsubscribeOption) error
	// Disconnect allows to close all connections of user to server.
	Disconnect(ctx context.Context, user string, opts ...gocent.DisconnectOption) error
	// Presence returns channel presence information.
	Presence(ctx context.Context, channel string) (gocent.PresenceResult, error)
	// PresenceStats returns short channel presence information (only counters).
	PresenceStats(ctx context.Context, channel string) (gocent.PresenceStatsResult, error)
	// History returns channel history.
	History(ctx context.Context, channel string, opts ...gocent.HistoryOption) (gocent.HistoryResult, error)
	// HistoryRemove removes channel history.
	HistoryRemove(ctx context.Context, channel string) error
	// Channels returns information about active channels (with one or more subscribers) on server.
	Channels(ctx context.Context, opts ...gocent.ChannelsOption) (gocent.ChannelsResult, error)
	// Info returns information about server nodes.
	Info(ctx context.Context) (gocent.InfoResult, error)
	// SendPipe sends Commands collected in Pipe to Centrifugo. Using this method you
	// should manually inspect all replies.
	SendPipe(ctx context.Context, pipe *gocent.Pipe) ([]gocent.Reply, error)
}
