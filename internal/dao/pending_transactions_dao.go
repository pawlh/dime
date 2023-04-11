package dao

import (
	"dime/internal/models"
)

type Pending interface {
	Create(transactions *models.PendingTransactions) error
	FindByOwner(owner string) (*models.PendingTransactions, error)
}
