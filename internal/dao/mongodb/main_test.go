package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var client *mongo.Client

var testURI = "mongodb://localhost:27018"

func newClient() (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(testURI))
	if err != nil {
		return nil, err
	}
	err = client.Connect(nil)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func BeforeEach() {
	var err error
	client, err = newClient()
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	err = client.Database("dime").Drop(nil)
	if err != nil {
		log.Fatalf("Error dropping database: %v", err)
	}
}
