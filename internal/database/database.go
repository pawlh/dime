package database

import "dime/internal/database/dao"

type Database interface {
	UserDAO() (dao.UserDAO, error)
	Disconnect()
}
