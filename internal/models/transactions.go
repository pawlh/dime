package models

// Transactions can hold any number of columns, but must at least have the following columns:
// ID          int
// Date        time.Time
// Description string
// Amount      float64
// Category    string
// Account     string
type Transactions struct {
	Transactions []map[string]any `bson:"transactions" json:"transactions"`
	Owner        string           `bson:"owner" json:"owner"`
	Columns      []string         `bson:"columns" json:"columns"`
}

var RequiredColumns = map[string]string{
	"id":          "string",
	"date":        "date",
	"description": "string",
	"amount":      "float",
	"category":    "string",
	"account":     "string",
}
