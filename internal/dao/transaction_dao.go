package dao

import (
	"dime/internal/models"
)

type TransactionDao interface {
	Insert(transaction *models.Transaction) error
}
