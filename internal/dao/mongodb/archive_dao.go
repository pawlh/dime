package mongodb

import (
	"dime/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Archive struct {
	client *mongo.Client
}

func NewArchive(client *mongo.Client) Archive {
	return Archive{client: client}
}

func (m Archive) Create(archive *models.Archive) (string, error) {

	// the ID is set by the database
	archive.ID = ""

	collections := m.client.Database("dime").Collection("archive")
	result, err := collections.InsertOne(nil, archive)
	if err != nil {
		return "", err
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
}

func (m Archive) UpdateColumnMapping(id string, columnMapping *models.ColumnMapping) error {
	collection := m.client.Database("dime").Collection("archive")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	filter := bson.D{{"_id", objectID}}
	update := bson.D{{"$set", bson.D{{"column_mapping", columnMapping}}}}

	_, err = collection.UpdateOne(nil, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (m Archive) FindByID(id string) (*models.Archive, error) {
	collection := m.client.Database("dime").Collection("archive")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{"_id", objectID}}
	var archive models.Archive
	err = collection.FindOne(nil, filter).Decode(&archive)
	if err != nil {
		return nil, err
	}

	return &archive, nil
}

func (m Archive) FindByOwner(owner string) ([]*models.Archive, error) {
	collection := m.client.Database("dime").Collection("archive")

	filter := bson.D{{"owner", owner}}
	var archives []*models.Archive
	cursor, err := collection.Find(nil, filter)
	if err != nil {
		return nil, err
	}

	for cursor.Next(nil) {
		var archive models.Archive
		err = cursor.Decode(&archive)
		if err != nil {
			return nil, err
		}

		archives = append(archives, &archive)
	}

	return archives, nil
}
