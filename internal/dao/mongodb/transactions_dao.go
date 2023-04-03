package mongodb

import (
	"dime/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Transactions struct {
	client *mongo.Client
}

func NewTransactions(client *mongo.Client) Transactions {
	return Transactions{client: client}
}

func (m Transactions) Insert(transaction *models.Transactions) error {
	// if a transaction with the same Owner already exists, just add the transactions to the existing one
	collection := m.client.Database("dime").Collection("transactions")

	var existingTransactions models.Transactions
	filter := bson.D{{"owner", transaction.Owner}}
	err := collection.FindOne(nil, filter).Decode(&existingTransactions)
	if err == mongo.ErrNoDocuments {
		// no existing transaction found, so insert the new one
		_, err := collection.InsertOne(nil, transaction)
		if err != nil {
			return err
		}

		return nil
	} else if err != nil {
		return err
	}

	// existing transaction found, so update it

	filter = bson.D{{"owner", transaction.Owner}}
	update := bson.D{{"$push", bson.D{{"transactions", bson.D{{"$each", transaction.Transactions}}}}}}

	_, err = collection.UpdateOne(nil, filter, update)
	if err != nil {
		return err
	}

	return nil

}

func (m Transactions) FindByOwner(owner string) (*models.Transactions, error) {
	collection := m.client.Database("dime").Collection("transactions")

	filter := bson.D{{"owner", owner}}
	var transactions models.Transactions
	err := collection.FindOne(nil, filter).Decode(&transactions)
	if err == mongo.ErrNoDocuments {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &transactions, nil
}
