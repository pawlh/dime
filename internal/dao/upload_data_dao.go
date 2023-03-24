package dao

import "dime/internal/models"

type UploadDataDAO interface {
	Create(uploadData *models.UploadData) error
}
