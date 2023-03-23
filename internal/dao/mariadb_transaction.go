package dao

import (
	"database/sql"
	"dime/internal/models"
)

type MariaDbTransaction struct {
	db *sql.DB
}

func NewMariaDbTransaction(db *sql.DB) MariaDbTransaction {
	return MariaDbTransaction{db: db}
}

func (m MariaDbTransaction) Insert(transaction *models.Transaction) error {
	_, err := m.db.Exec("INSERT INTO transaction (date, description, amount, category, account) VALUES (?, ?, ?, ?, ?)",
		transaction.Date,
		transaction.Description,
		transaction.Amount,
		transaction.Category,
		transaction.Account,
	)
	return err
}
