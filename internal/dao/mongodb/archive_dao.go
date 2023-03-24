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

	return nil
}
