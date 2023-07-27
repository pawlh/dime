package mongodb

import (
	"dime/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const transactionCollectionName = "transaction"

type TransactionDAO struct {
	client *mongo.Client
}

func NewTransactionDAO(client *mongo.Client) *TransactionDAO {
	return &TransactionDAO{client: client}
}

func (dao TransactionDAO) AddTransaction(transaction models.Transaction) (string, error) {
	collection := dao.client.Database("dime").Collection(transactionCollectionName)
	result, err := collection.InsertOne(nil, transaction)
	if err != nil {
		return "", err
	}

	return objectedIdToHex(result.InsertedID.(primitive.ObjectID)), nil
}

func (dao TransactionDAO) AddTransactions(transactions []models.Transaction) error {

	newTransactions := make([]interface{}, len(transactions))
	for i, transaction := range transactions {
		newTransactions[i] = transaction
	}

	collection := dao.client.Database("dime").Collection(transactionCollectionName)
	_, err := collection.InsertMany(nil, newTransactions)
	if err != nil {
		return err
	}

	return nil
}

func (dao TransactionDAO) GetTransactions(owner string) ([]models.Transaction, error) {
	collection := dao.client.Database("dime").Collection(transactionCollectionName)
	filter := bson.D{{"owner", owner}}

	var transactions []models.Transaction
	cursor, err := collection.Find(nil, filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(nil, &transactions)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func (dao TransactionDAO) Clear() error {
	collection := dao.client.Database("dime").Collection(transactionCollectionName)
	_, err := collection.DeleteMany(nil, bson.D{})
	return err
}
