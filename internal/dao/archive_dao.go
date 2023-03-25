package dao

import "dime/internal/models"

type ArchiveDao interface {
	Create(archive *models.Archive) error
	UpdateColumnMapping(archive *models.Archive) error
}
