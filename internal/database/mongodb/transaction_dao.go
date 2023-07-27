package mongodb

import "go.mongodb.org/mongo-driver/mongo"

const transactionCollectionName = "transaction"

type TransactionDAO struct {
	client *mongo.Client
}

func (dao TransactionDAO) AddTransaction(owner string, transaction []map[string]any) (string, error) {
	panic("implement me")
}

func (dao TransactionDAO) AddTransactions(owner string, transactions []map[string]any) (string, error) {
	panic("implement me")
}

func (dao TransactionDAO) GetTransactions(owner string) ([]map[string]any, error) {
	panic("implement me")
}
