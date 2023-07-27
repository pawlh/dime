package models

type Transaction struct {
	Id          string           `bson:"_id,omitempty"`
	Owner       string           `bson:"owner"`
	Transaction []map[string]any `bson:"transaction"`
}

// RequiredColumns is a map of the base required columns for a transaction
var RequiredColumns = map[string]string{
	"id":          "string",
	"date":        "date",
	"description": "string",
	"amount":      "float",
	"category":    "string",
	"account":     "string",
}
