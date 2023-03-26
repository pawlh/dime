package dao

import "dime/internal/models"

type ArchiveDao interface {
	Create(archive *models.Archive) (string, error)
	UpdateColumnMapping(archive *models.Archive) error
	FindByID(id string) (*models.Archive, error)
}
