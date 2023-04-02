package dao

import (
	"dime/internal/models"
)

type TransactionsDao interface {
	Insert(transactions *models.Transactions) error
}
