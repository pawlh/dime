package dao

import (
	"database/sql"
	"dime/internal/models"
)

type MariaDbUser struct {
	db *sql.DB
}

func NewMariaDbUser(db *sql.DB) MariaDbUser {
	return MariaDbUser{db: db}
}

func (m MariaDbUser) Insert(user *models.User) error {
	_, err := m.db.Exec("INSERT INTO user (username, password) VALUES (?, ?)",
		user.Username,
		user.Password,
	)
	return err
}
