package mariadb

import (
	"database/sql"
	"dime/internal/models"
)

type UploadData struct {
	db *sql.DB
}

func NewMariaDbUploadData(db *sql.DB) UploadData {
	return UploadData{db: db}
}

func (m UploadData) Create(uploadData *models.UploadData) error {
	_, err := m.db.Exec("INSERT INTO upload_data (upload_date, file_name, original_name, owner) VALUES (?, ?, ?, ?)",
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

func (m UploadData) FindByOwner(owner string) ([]models.UploadData, error) {
	rows, err := m.db.Query("SELECT * FROM upload_data WHERE owner = ?", owner)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) error {
		err := rows.Close()
		if err != nil {
			return err
		}
		return nil
	}(rows)

	var uploadData []models.UploadData
	for rows.Next() {
		var ud models.UploadData
		err := rows.Scan(&ud.ID, &ud.UploadDate, &ud.FileName, &ud.OriginalName, &ud.Owner)
		if err != nil {
			return nil, err
		}
		uploadData = append(uploadData, ud)
	}
	return uploadData, nil
}
