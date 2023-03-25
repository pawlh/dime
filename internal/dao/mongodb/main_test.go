package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testing"
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

func TestMain(m *testing.M) {
	var err error
	client, err = newClient()
	if err != nil {
		panic(err)
	}
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(client, nil)

	m.Run()

	err = client.Database("dime").Drop(nil)
	if err != nil {
		log.Fatalf("Error dropping database: %v", err)
	}
}
