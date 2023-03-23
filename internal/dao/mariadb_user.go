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

func (m MariaDbUser) Create(user *models.User) error {
	_, err := m.db.Exec("INSERT INTO user (username, password, name) VALUES (?, ?, ?)",
		user.Username,
		user.Password,
		user.Name,
	)
	return err
}

func (m MariaDbUser) FindByUsername(username string) (*models.User, error) {
	row := m.db.QueryRow("SELECT username, password, name FROM user WHERE username = ?", username)
	user := new(models.User)
	err := row.Scan(&user.Username, &user.Password)
	if err == sql.ErrNoRows {
		return nil, UserNotFound{}
	} else if err != nil {
		return nil, err
	}
	return user, nil
}
