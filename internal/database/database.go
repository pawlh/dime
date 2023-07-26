package database

import "dime/internal/database/dao"

type Database interface {
	// UserDAO returns a dao.UserDAO instance. Init() must be called before this method
	UserDAO() (dao.UserDAO, error)
	Init(string)

	// Disconnect closes the connection to the database. This method should be called when no more database operations are needed
	Disconnect()
}
