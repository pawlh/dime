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

func (m MariaDbUser) FindByUsername(username string) (*models.User, error) {
	row := m.db.QueryRow("SELECT username, password FROM user WHERE username = ?", username)
	user := new(models.User)
	err := row.Scan(&user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
