package models

import "time"

type Archive struct {
	ID           int
	UploadDate   time.Time `bson:"upload_date"`
	OriginalName string    `bson:"original_name"`
	Data         []map[string]string
}
