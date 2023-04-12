package mongodb

import (
	"dime/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PendingTransactions struct {
	client *mongo.Client
}

func NewPendingTransactions(client *mongo.Client) PendingTransactions {
	return PendingTransactions{client: client}
}

// Create saves a new pending transaction to the database
func (m PendingTransactions) Create(transactions *models.PendingTransactions) error {
	collection := m.client.Database("dime").Collection("pending_transactions")

	_, err := collection.InsertOne(nil, transactions)
	if err != nil {
		return err
	}

	return nil
}

func (m PendingTransactions) FindByOwner(owner string) (*models.PendingTransactions, error) {
	collection := m.client.Database("dime").Collection("transactions")

	var transactions models.PendingTransactions
	err := collection.FindOne(nil, bson.M{"owner": owner}).Decode(&transactions)
	if err != nil {
		return nil, err
	}

	return &transactions, nil
}
