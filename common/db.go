package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseConnection struct {
	uri    string
	client *mongo.Client
	db     *mongo.Database
	// todo: add context
}

func initDBConn(uri string, dbName string) (DatabaseConnection, error) {
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	db := client.Database(dbName)

	if err != nil {
		return DatabaseConnection{}, err
	}

	return DatabaseConnection{
		uri,
		client,
		db,
	}, nil
}

func main() {
	conn, err := initDBConn("mongodb://root:test@127.0.0.1:27017", "test-db")

	var result bson.M
	err = conn.db.RunCommand(context.TODO(), bson.D{{"dbStats", 1}}).Decode(&result)

	if err == nil {
		fmt.Println(result)
	} else {
		panic(err)
	}
}
