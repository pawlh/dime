package database

import "dime/internal/database/dao"

type Database interface {
	UserDAO() (dao.UserDAO, error)
	TransactionDAO() (dao.TransactionDAO, error)

	// Disconnect closes the connection to the database. This method should be called when no more database operations are needed
	Disconnect()
}
