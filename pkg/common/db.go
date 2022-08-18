package common

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

type MongoConn struct {
	coll *mongo.Collection
}

func (m *MongoConn) InitConn() error {
	uri := os.Getenv("MONGODB_URI")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}

	m.coll = client.Database("housing").Collection("housing")
	return nil
}

func (m MongoConn) Insert(listing Listing) (*mongo.InsertOneResult, error) {
	res, err := m.coll.InsertOne(context.TODO(), listing)
	return res, err
}

func (m MongoConn) RetrieveRecords(count int64) ([]Listing, error) {
	// retrieve count records
	findOptions := options.Find()
	findOptions.SetLimit(count)
	var results []Listing
	res, err := m.coll.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		return []Listing{}, err
	}

	for res.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var elem Listing
		err := res.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)
	}

	return results, nil
}

func (m MongoConn) Exists(query bson.D) bool {
	var result bson.M
	return m.coll.FindOne(context.TODO(), query).Decode(&result) != nil
}
