package dao

import "dime/internal/models"

type UserDAO interface {
	Insert(user *models.User) error
}
