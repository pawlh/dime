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

// Insert This is a nasty non-mongodb native solution to the problem of upserting a slice of transactions
func (m Transactions) Insert(transactions *models.Transactions) error {
	// if a transactions with the same Owner already exists, just add the transactions to the existing one
	collection := m.client.Database("dime").Collection("transactions")

	var existingTransactions models.Transactions
	filter := bson.D{{"owner", transactions.Owner}}
	err := collection.FindOne(nil, filter).Decode(&existingTransactions)
	if err == mongo.ErrNoDocuments {
		// no existing transactions found, so insert the new one
		_, err := collection.InsertOne(nil, transactions)
		if err != nil {
			return err
		}

		return nil
	} else if err != nil {
		return err
	}

	for _, transaction := range transactions.Transactions {
		filter := bson.D{{"owner", transactions.Owner}, {"transactions.id", transaction["id"]}}
		count, err := collection.CountDocuments(nil, filter)
		if err != nil {
			return err
		} else if count != 0 {
			continue
		}

		update := bson.D{{"$push", bson.D{{"transactions", transaction}}}}
		filter = bson.D{{"owner", transactions.Owner}}
		_, err = collection.UpdateOne(nil, filter, update)
		if err != nil {
			return err
		}
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

func contains(s []interface{}, e interface{}) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
