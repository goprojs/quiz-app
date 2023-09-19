package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	username := os.Getenv("MONGODB_USERNAME")
	password := os.Getenv("MONGODB_PASSWORD")
	mongo_uri := fmt.Sprintf("mongodb+srv://%s:%s@cluster0.vk2ln2m.mongodb.net/?retryWrites=true&w=majority", username, password)

	// Connect
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongo_uri))
	if err != nil {
		log.Fatal(err)
	}

	// Disconnect
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// Ping
	// err = client.Ping(ctx, readpref.Primary())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Add databases
	quickStartDatabase := client.Database("quickstart")
	bookCollections := quickStartDatabase.Collection("book_collections")
	//copyCollections := quickStartDatabase.Collection("copy_collections")
	bookResult, err := bookCollections.InsertOne(ctx, bson.D{
		{"title", "The King"},
		{"author", "Chetan Bhagat"},
		{"tags", bson.A{"fiction", "short story"}},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bookResult.InsertedID)

	// Databases
	databases, err := client.ListDatabaseNames(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
}
