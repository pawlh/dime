package mongodb

import (
	"dime/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const userCollectionName = "user"

type UserDAO struct {
	client *mongo.Client
}

func NewUserDAO(client *mongo.Client) *UserDAO {
	return &UserDAO{client: client}
}

func (dao UserDAO) AddUser(user models.User) (string, error) {
	collection := dao.client.Database("dime").Collection(userCollectionName)
	result, err := collection.InsertOne(nil, user)
	if err != nil {
		return "", err
	}

	return objectedIdToHex(result.InsertedID.(primitive.ObjectID)), nil
}

func (dao UserDAO) GetUser(id string) (*models.User, error) {
	collection := dao.client.Database("dime").Collection(userCollectionName)
	filter := bson.D{{"_id", hexToObjectId(id)}}

	var user models.User
	err := collection.FindOne(nil, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else {
			return nil, err
		}
	}

	return &user, nil
}

// TODO: Move this to a utils package
// objectedIdToHex converts an objectId to a hex string. Note, must cast the objectId to a primitive.ObjectID
func objectedIdToHex(objectId primitive.ObjectID) string {
	return objectId.Hex()
}

// TODO: Move this to a utils package
// hexToObjectId converts a hex string to an objectId
func hexToObjectId(hex string) primitive.ObjectID {
	objectId, _ := primitive.ObjectIDFromHex(hex)
	return objectId
}
