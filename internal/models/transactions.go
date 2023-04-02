package models

// Transactions can hold any number of columns, but must at least have the following columns:
// ID          int
// Date        time.Time
// Description string
// Amount      float64
// Category    string
// Account     string
type Transactions struct {
	Transactions []map[string]interface{} `bson:"transactions"`
	Owner        string                   `bson:"owner"`
	Columns      []string                 `bson:"columns"`
}
