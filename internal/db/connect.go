package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connect() (*mongo.Client, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	opts := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Printf("connection error :%v", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		log.Println(err)
	}

	return client, cancel
}

func GetCollection(collectionName string) (*mongo.Collection, context.CancelFunc) {
	client, cancel := connect()
	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection(collectionName)

	return collection, cancel
}
