package models

import "time"

type Archive struct {
	ID           int
	UploadDate   time.Time
	FileName     string
	OriginalName string
	Owner        string
}
