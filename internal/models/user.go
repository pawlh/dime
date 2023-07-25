package models

type User struct {
	FirstName string `bson:"firstName"`
	LastName  string `bson:"lastName"`
	Password  string `bson:"password"`
}
