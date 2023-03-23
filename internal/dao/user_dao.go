package dao

import "dime/internal/models"

type UserNotFound struct{}

func (e UserNotFound) Error() string {
	return "user not found"
}

type UserDAO interface {
	Create(user *models.User) error
	FindByUsername(username string) (*models.User, error)
}
