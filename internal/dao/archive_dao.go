package dao

import "dime/internal/models"

type ArchiveDao interface {
	Create(archive *models.Archive) (string, error)
	UpdateColumnMapping(string, *models.ColumnMapping) error
	FindByID(id string) (*models.Archive, error)
	FindByOwner(owner string) ([]*models.Archive, error)
}
