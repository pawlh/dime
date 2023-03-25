package mongodb

import (
	"dime/internal/models"
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
