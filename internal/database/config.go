package database

import (
	"dime/internal/database/mongodb"
	"os"
)

var DB Database = mongodb.Init(os.Getenv("MONGODB_URI"))
