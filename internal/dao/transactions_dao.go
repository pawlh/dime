package dao

import (
	"dime/internal/models"
)

type TransactionsDao interface {
	Insert(transactions *models.Transactions) error
	FindByOwner(owner string) (*models.Transactions, error)
}
