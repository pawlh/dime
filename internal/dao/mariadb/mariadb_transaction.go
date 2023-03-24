package mariadb

import (
	"database/sql"
	"dime/internal/models"
)

type Transaction struct {
	db *sql.DB
}

func NewMariaDbTransaction(db *sql.DB) Transaction {
	return Transaction{db: db}
}

func (m Transaction) Insert(transaction *models.Transaction) error {
	_, err := m.db.Exec("INSERT INTO transaction (date, description, amount, category, account) VALUES (?, ?, ?, ?, ?)",
		transaction.Date,
		transaction.Description,
		transaction.Amount,
		transaction.Category,
		transaction.Account,
	)
	return err
}
