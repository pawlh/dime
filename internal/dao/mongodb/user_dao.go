package mongodb

import (
	"dime/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserDao struct {
	client *mongo.Client
}

func NewUser(client *mongo.Client) UserDao {
	return UserDao{client: client}
}

func (m UserDao) Create(user *models.User) error {
	collection := m.client.Database("dime").Collection("user")
	_, err := collection.InsertOne(nil, user)

	return err
}

func (m UserDao) FindByUsername(username string) (*models.User, error) {
	collection := m.client.Database("dime").Collection("user")
	filter := bson.M{"username": username}

	var user models.User
	err := collection.FindOne(nil, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
