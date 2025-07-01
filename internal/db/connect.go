package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*mongo.Client, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	opts := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatalf("connection error :%v", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}

	return client, cancel
}

func GetCollection(client *mongo.Client, collection string) *mongo.Collection {
	return client.Database(os.Getenv("MONGO_DATABASE")).Collection(collection)
}
