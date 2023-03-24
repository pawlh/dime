package api

import (
	"bytes"
	"dime/internal/dbs"
	"dime/internal/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

var dateLayout = "2006-01-02 15:04:05"

func Upload(c echo.Context) error {

	formFile, err := c.FormFile("file")
	if err != nil {
		return mustSendError(c, http.StatusBadRequest, "missing file", err)
	}

	file, err := formFile.Open()
	if err != nil {
		return mustSendError(c, http.StatusInternalServerError, "error opening file", err)
	}

	uploadDate := time.Now()

	fileName := c.Get("username").(string) + "_" + uploadDate.Format(dateLayout) + ".csv"
	archive := models.Archive{
		UploadDate:   uploadDate,
		Owner:        c.Get("username").(string),
		FileName:     fileName,
		OriginalName: formFile.Filename,
	}

	err = saveCSVFile(file, fileName)
	if err != nil {
		return mustSendError(c, http.StatusInternalServerError, "error saving file", err)
	}

	err = dbs.DB.ArchiveDao().Create(&archive)
	if err != nil {
		return mustSendError(c, http.StatusInternalServerError, "error saving upload record", err)
	}

	return nil
}

func saveCSVFile(file multipart.File, fileName string) error {
	buf := new(bytes.Buffer)
	_, err := io.Copy(buf, file)
	if err != nil {
		return fmt.Errorf("error copying file: %w", err)
	}

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(path)

	_, err = os.Stat("storage/uploads")
	if os.IsNotExist(err) {
		err = os.MkdirAll("storage/uploads", 0755)
		if err != nil {
			return fmt.Errorf("error creating directory: %w", err)
		}
	}
	err = os.WriteFile("storage/uploads/"+fileName, buf.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("error saving file: %w", err)
	}

	return nil
}
