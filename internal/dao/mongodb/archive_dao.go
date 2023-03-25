package mongodb

import (
	"dime/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Archive struct {
	client *mongo.Client
}

func NewArchive(client *mongo.Client) Archive {
	return Archive{client: client}
}

func (m Archive) Create(archive *models.Archive) error {

	// the ID is set by the database
	archive.ID = ""

	collections := m.client.Database("dime").Collection("archive")
	_, err := collections.InsertOne(nil, archive)
	if err != nil {
		return err
	}

	return nil
}

func (m Archive) UpdateColumnMapping(archive *models.Archive) error {
	collection := m.client.Database("dime").Collection("archive")

	update := bson.D{{"$set", bson.D{{"column_mapping", archive.ColumnMapping}}}}

	_, err := collection.UpdateOne(nil, models.Archive{ID: archive.ID}, update)
	if err != nil {
		return err
	}

	return nil
}
