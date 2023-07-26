package models

type User struct {
	Id        string `bson:"_id,omitempty"`
	FirstName string `bson:"firstName"`
	LastName  string `bson:"lastName"`
	Password  string `bson:"password"`
}
