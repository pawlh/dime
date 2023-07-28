package models

import "time"

type Transaction struct {
	Id          string         `bson:"_id,omitempty" json:"id"`
	Owner       string         `bson:"owner" json:"owner"`
	BaseFields  BaseFields     `bson:"baseFields" json:"baseFields"`
	ExtraFields map[string]any `bson:"extraFields" json:"extraFields"`
}

type BaseFields struct {
	Date        time.Time `bson:"date" json:"date"`
	Description string    `bson:"description" json:"description"`
	Amount      float64   `bson:"amount" json:"amount"`
	Category    string    `bson:"category" json:"category"`
	Account     string    `bson:"account" json:"account"`
}
