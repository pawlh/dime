package models

import "time"

type Transaction struct {
	Id          string         `bson:"_id,omitempty"`
	Owner       string         `bson:"owner"`
	BaseFields  BaseFields     `bson:"baseFields"`
	ExtraFields map[string]any `bson:"extraFields"`
}

type BaseFields struct {
	Id          string    `bson:"id,omitempty"`
	Date        time.Time `bson:"date"`
	Description string    `bson:"description"`
	Amount      float64   `bson:"amount"`
	Category    string    `bson:"category"`
	Account     string    `bson:"account"`
}
