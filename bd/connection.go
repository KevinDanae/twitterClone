package bd

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN = MongoConnection() // MongoC is the object that represents the connection to the database

// MongoConection is a function that returns a mongo client and connects to the database
func MongoConnection() *mongo.Client {

	values := os.Getenv("MONGODB_URI")
	if values == "" {
		values = "mongodb+srv://KevinAdmin:TR2ntQfmNg1kmyGx@cluster0.wph5t.mongodb.net/twitterClone?retryWrites=true&w=majority"
	}

	clientOptions := options.Client().ApplyURI(values)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Connected to MongoDB!")
	return client
}

// CheckConnection is a function that checks if the connection to the database is active
func CheckConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
