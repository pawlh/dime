package dao

import "dime/internal/models"

type TransactionDAO interface {
	// AddTransaction Add a transaction to the database
	AddTransaction(transaction models.Transaction) (string, error)
	// AddTransactions Add a list of transactions to the database.
	AddTransactions(transactions []models.Transaction) error
	// GetTransactions Fetch all transactions for a user
	GetTransactions(owner string) ([]models.Transaction, error)
	Clear() error
}
