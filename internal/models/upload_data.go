package models

import "time"

type UploadData struct {
	ID           int
	UploadDate   time.Time
	FileName     string
	OriginalName string
	Owner        string
}
