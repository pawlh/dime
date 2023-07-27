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

func (dao UserDAO) Clear() error {
	collection := dao.client.Database("dime").Collection(userCollectionName)
	_, err := collection.DeleteMany(nil, bson.D{})
	return err
}
