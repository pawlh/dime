package dao

import "dime/internal/models"

type UserDAO interface {

	// GetUser Fetch a user from the database
	GetUser(id string) (*models.User, error)

	// AddUser Create a new user.
	// Returns the UUID of the new user
	AddUser(user models.User) (string, error)
}
