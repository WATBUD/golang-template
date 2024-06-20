package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewMongoClient() (*mongo.Client, error) {
	// connect mongodb
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Create a new client and connect to the server
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://williamchuang:hkdlOv3cRPlHUTR1@maitoday.kfzdeqh.mongodb.net/?retryWrites=true&w=majority&appName=MaiToday"))
	if err != nil {
		return nil, fmt.Errorf("connent to mongodb error: %s", err)
	}

	// Send a ping to confirm a successful connection
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, fmt.Errorf("connent to mongodb error: %s", err)
	}

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	return client, nil
}
