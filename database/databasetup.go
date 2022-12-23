package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBSet() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeOut(context.Background(), 10*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	client.Ping(context.TODO(), nil)
	if err != nil {
		log.Println("Failed to connect to MongoDB :(")
		return nil
	}

	fmt.Println("Successfully connectod to MongoDB :)")

	return client
}

var Client *mongo.Client = DBSet()

func UserData(client *mongo.Client, collectionName string) *mongo.Collection {
	return nil
}

func ProductData(client *mongo.Client, collectionName string) *mongo.Collection {
	return nil
}
