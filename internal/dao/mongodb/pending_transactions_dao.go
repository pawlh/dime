package mongodb

import (
	"dime/internal/models"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PendingTransactions struct {
	client *mongo.Client
}

func NewPendingTransactions(client *mongo.Client) PendingTransactions {
	return PendingTransactions{client: client}
}

// Create saves a new pending transaction to the database
func (m PendingTransactions) Create(transactions *models.PendingTransactions) (string, error) {

	if transactions.Owner == "" {
		return "", errors.New("owner is required")
	}
	if transactions.Name == "" {
		return "", errors.New("name is required")
	}
	if transactions.WIPTransactions == nil || len(transactions.WIPTransactions) == 0 {
		return "", errors.New("wip_transactions is required")
	}

	collection := m.client.Database("dime").Collection("pending_transactions")

	result, err := collection.InsertOne(nil, transactions)
	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (m PendingTransactions) FindByOwner(owner string) (*models.PendingTransactions, error) {
	collection := m.client.Database("dime").Collection("pending_transactions")

	var transactions models.PendingTransactions
	err := collection.FindOne(nil, bson.M{"owner": owner}).Decode(&transactions)
	if err != nil {
		return nil, err
	}

	return &transactions, nil
}

func (m PendingTransactions) FindById(id string) (*models.PendingTransactions, error) {
	collection := m.client.Database("dime").Collection("pending_transactions")

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var transactions models.PendingTransactions
	err = collection.FindOne(nil, bson.M{"_id": objectId}).Decode(&transactions)
	if err != nil {
		return nil, err
	}

	return &transactions, nil
}

func (m PendingTransactions) Save(transactions *models.PendingTransactions) error {
	collection := m.client.Database("dime").Collection("pending_transactions")

	objectId, err := primitive.ObjectIDFromHex(transactions.TransactionGroupId)
	if err != nil {
		return err
	}

	//set "transactions" to match wip_transactions

	_, err = collection.UpdateOne(nil,
		bson.M{"_id": objectId},
		bson.M{"$set": bson.M{"transactions": transactions.WIPTransactions}},
	)
	//_, err = collection.ReplaceOne(nil, bson.M{"_id": objectId}, transactions)
	if err != nil {
		return err
	}

	return nil
}

func (m PendingTransactions) Clear() error {
	collection := m.client.Database("dime").Collection("pending_transactions")
	err := collection.Drop(nil)
	if err != nil {
		return err
	}

	return nil
}
