package mongodb

import (
	"context"
	"dime/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"reflect"
	"testing"
)

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

var client *mongo.Client

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

func TestUserDao_Create_Find(t *testing.T) {
	testUser := models.User{
		Username: "testUsername",
		Password: "testPassword",
		Name:     "testName",
	}

	userDao := NewUser(client)
	err := userDao.Create(&testUser)
	if err != nil {
		t.Errorf("Error creating a new user: %v", err)
	}

	if match, err := userDao.FindByUsername(testUser.Username); err != nil {
		t.Errorf("Error finding user: %v", err)
	} else if !reflect.DeepEqual(testUser, *match) {
		t.Errorf("Users do not match: %v", err)
	}
}
