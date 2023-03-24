package dbs

import "dime/internal/dao"

type Database interface {
	TransactionDao() dao.TransactionDao
	UserDao() dao.UserDAO
	UploadDataDao() dao.UploadDataDAO
}

var DB Database
