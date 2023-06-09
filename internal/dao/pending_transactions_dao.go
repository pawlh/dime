package dao

import (
	"dime/internal/models"
)

type PendingTransactionsDao interface {
	Create(transactions *models.PendingTransactions) (string, error)
	FindByOwner(owner string) (*models.PendingTransactions, error)
	FindById(id string) (*models.PendingTransactions, error)
	Save(transactions *models.PendingTransactions) error
	Clear() error
}
