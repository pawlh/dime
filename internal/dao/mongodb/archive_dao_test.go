package mongodb

import (
	"dime/internal/models"
	"testing"
	"time"
)

func TestArchive_Create(t *testing.T) {
	testArchive := models.Archive{
		ID:           0,
		UploadDate:   time.Time{},
		OriginalName: "",
		Data:         nil,
	}

	archiveDao := NewArchive(client)
	err := archiveDao.Create(&testArchive)
	if err != nil {
		t.Errorf("Error creating a new archive: %v", err)
	}
}
