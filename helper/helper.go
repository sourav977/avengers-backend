package helper

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//ConnectToDB connects to db, creates a collection and return it
func ConnectToDB() *mongo.Collection {
	uri := os.Getenv("MONGO_CONNECTION_URL")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	// defer func() {
	// 	if err = client.Disconnect(ctx); err != nil {
	// 		panic(err)
	// 	}
	// }()

	// connection test
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
		panic(err)
	}
	fmt.Println("Successfully connected and pinged.")
	collection := client.Database("my-db").Collection("avengers")
	return collection
}
