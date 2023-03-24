package mariadb

import (
	"database/sql"
	"dime/internal/models"
	"github.com/labstack/gommon/log"
)

type UploadData struct {
	db *sql.DB
}

func NewMariaDbArchive(db *sql.DB) UploadData {
	return UploadData{db: db}
}

func (m UploadData) Create(uploadData *models.Archive) error {
	_, err := m.db.Exec("INSERT INTO archive (upload_date, file_name, original_name, owner) VALUES (?, ?, ?, ?)",
		uploadData.UploadDate,
		uploadData.FileName,
		uploadData.OriginalName,
		uploadData.Owner,
	)
	if err != nil {
		return err
	}

	return nil
}

func (m UploadData) FindByOwner(owner string) ([]models.Archive, error) {
	rows, err := m.db.Query("SELECT * FROM archive WHERE owner = ?", owner)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Error(err)
		}
	}(rows)

	var uploadData []models.Archive
	for rows.Next() {
		var ud models.Archive
		err := rows.Scan(&ud.ID, &ud.UploadDate, &ud.FileName, &ud.OriginalName, &ud.Owner)
		if err != nil {
			return nil, err
		}
		uploadData = append(uploadData, ud)
	}
	return uploadData, nil
}
