package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connection URI
const uri = "mongodb://user:pass@sample.host:27017/?maxPoolSize=20&w=majority"

type DatabaseConnection struct {
	uri    string
	client *mongo.Client
}

func initDBConn(uri string) (DatabaseConnection, error) {
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err == nil {
		return DatabaseConnection{}, err
	}

	dbConn := DatabaseConnection{
		uri,
		client,
	}

	return dbConn, nil
}

func main() {
	fmt.Println("Successfully connected and pinged.")
}
