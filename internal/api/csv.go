package api

import (
	"bytes"
	"dime/internal/csv"
	"dime/internal/dbs"
	"dime/internal/models"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"mime/multipart"
	"net/http"
	"time"
)

//var dateLayout = "2006-01-02 15:04:05"

func Upload(c echo.Context) error {

	formFile, err := c.FormFile("file")
	if err != nil {
		return mustSendError(c, http.StatusBadRequest, "missing file", err)
	}

	file, err := formFile.Open()
	if err != nil {
		return mustSendError(c, http.StatusInternalServerError, "error opening file", err)
	}

	buf, err := fileToBuffer(file)
	if err != nil {
		return mustSendError(c, http.StatusInternalServerError, "error reading file", err)
	}

	parsedCSV, err := csv.Parse(buf)
	if err != nil {
		return mustSendError(c, http.StatusInternalServerError, "error parsing file", err)
	}

	uploadDate := time.Now()

	archive := models.Archive{
		UploadDate:   uploadDate,
		OriginalName: formFile.Filename,
		Owner:        c.Get("username").(string),
		Data:         parsedCSV,
	}

	_, err = dbs.DB.ArchiveDao().Create(&archive)
	if err != nil {
		return mustSendError(c, http.StatusInternalServerError, "error saving record", err)
	}

	return nil
}

func GetArchive(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return mustSendError(c, http.StatusBadRequest, "missing id", nil)
	}

	archive, err := dbs.DB.ArchiveDao().FindByID(id)
	if err != nil {
		return mustSendError(c, http.StatusInternalServerError, "error getting archive", err)
	}
	if archive == nil {
		return mustSendError(c, http.StatusNotFound, "archive not found", nil)
	}
	if archive.Owner != c.Get("username").(string) {
		return mustSendError(c, http.StatusForbidden, "not authorized", nil)
	}

	return c.JSON(http.StatusOK, archive)
}

func GetArchives(c echo.Context) error {
	archives, err := dbs.DB.ArchiveDao().FindByOwner(c.Get("username").(string))
	if err != nil {
		return mustSendError(c, http.StatusInternalServerError, "error getting archives", err)
	}

	return c.JSON(http.StatusOK, archives)
}

func fileToBuffer(file multipart.File) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	_, err := io.Copy(buf, file)
	if err != nil {
		return nil, fmt.Errorf("error copying file: %w", err)
	}

	return buf, nil
}
