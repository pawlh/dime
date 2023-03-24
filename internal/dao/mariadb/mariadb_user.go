package mariadb

import (
	"database/sql"
	"dime/internal/dao"
	"dime/internal/models"
)

type User struct {
	db *sql.DB
}

func NewMariaDbUser(db *sql.DB) User {
	return User{db: db}
}

func (m User) Create(user *models.User) error {
	_, err := m.db.Exec("INSERT INTO user (username, password, name) VALUES (?, ?, ?)",
		user.Username,
		user.Password,
		user.Name,
	)
	return err
}

func (m User) FindByUsername(username string) (*models.User, error) {
	row := m.db.QueryRow("SELECT username, password, name FROM user WHERE username = ?", username)
	user := new(models.User)
	err := row.Scan(&user.Username, &user.Password, &user.Name)
	if err == sql.ErrNoRows {
		return nil, dao.UserNotFound{}
	} else if err != nil {
		return nil, err
	}
	return user, nil
}
