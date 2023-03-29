package models

import "time"

// ColumnMapping provides information about how columns in Archive.Data map to models.Transaction
type ColumnMapping struct {
	Date        string
	Description string
	Amount      string
	Category    string
	Account     string
}

type Archive struct {
	ID            string        `bson:"_id,omitempty"`
	UploadDate    time.Time     `bson:"upload_date"`
	OriginalName  string        `bson:"original_name"`
	Owner         string        `bson:"owner"`
	ColumnMapping ColumnMapping `bson:"column_mapping"`
	Data          []map[string]string
}
