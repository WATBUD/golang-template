package mongodb

import (
	"context"
	"fmt"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	once sync.Once

	instance *mongo.Client
)

func Instance() *mongo.Client {
	once.Do(func() {
		instance = newClient()
	})
	return instance
}

func newClient() *mongo.Client {
	uri := defaultUriOrEnv()

	bsonOpts := &options.BSONOptions{
		UseJSONStructTags: true,
	}

	// connect mongodb
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetBSONOptions(bsonOpts))
	if err != nil {
		panic(fmt.Errorf("connent to mongodb error: %s", err))
	}

	// Send a ping to confirm a successful connection
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(fmt.Errorf("connent to mongodb error: %s", err))
	}

	return client
}

// CreateDatabase defines a function type that takes a pointer to a mongo.Database instance
// and returns a value of any type `T`. This allows for flexibility in what the function
// can create or return based on the database interaction.
type CreateDatabase[T any] func(*mongo.Database) T

// DefaultDatabaseClient defines a generic client for interacting with a default MongoDB database.
// It embeds a mongo.Client instance and provides a `CreateDatabase` function for
// customizing database interactions. The `T` type parameter allows the client to return
// different types based on the specific database operation.
type DefaultDatabaseClient[T any] struct {
	// Client embeds the underlying mongo.Client instance.
	*mongo.Client
	// CreateDatabase defines a callback function for creating or retrieving data from the database.
	// It takes a pointer to a mongo.Database and returns a value of type `T`.
	CreateDatabase CreateDatabase[T]
}

// Database returns a value of type `T` representing the result of interacting with the
// default database. It retrieves the database name and uses the `CreateDatabase` callback
// to perform the actual database interaction.
func (c *DefaultDatabaseClient[T]) Database(opts ...*options.DatabaseOptions) T {
	n := defaultDbNameOrEnv()
	db := c.Client.Database(n, opts...)
	return c.CreateDatabase(db)
}
