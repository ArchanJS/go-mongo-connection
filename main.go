package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

func main() {
	uri := "mongodb://localhost:27017"
	client, ctx, cancel, err := connect(uri)

	if err != nil {
		panic(err)
	}

	defer close(client, ctx, cancel)

	collection := client.Database("rupeek-hack").Collection("users")
	filter := bson.D{{"email", "archanbanerjee70@gmail.com"}}
	result, err := collection.Find(ctx, filter)
	if err != nil {
		panic(err)
	}
	// fmt.Println(result)
	for result.Next(ctx) {
		var document bson.M
		if err := result.Decode(&document); err != nil {
			panic(err)
		}
		fmt.Println(document)
	}
}
