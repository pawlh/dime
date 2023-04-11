package dao

import (
	"dime/internal/models"
)

type PendingTransactionsDao interface {
	Create(transactions *models.PendingTransactions) error
	FindByOwner(owner string) (*models.PendingTransactions, error)
}
