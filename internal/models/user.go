package models

type User struct {
	Id        string `bson:"_id,omitempty" json:"id"`
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"lastName" json:"lastName"`
	Password  string `bson:"password" json:"password,omitempty"`
}
