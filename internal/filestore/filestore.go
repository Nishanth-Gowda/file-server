package filestore

import (
	"errors"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nishanth-gowda/file-server/internal/repository"
)

type FileStore struct {
	repo repository.Repository
}

func NewFileStore(repo repository.Repository) *FileStore {
	return &FileStore{
		repo: repo,
	}
}

func (fs *FileStore) HandleStore(c *gin.Context) {
	fileName := c.Query("fileName")
	if fileName == "" {
		c.String(http.StatusBadRequest, "fileName is required")
		return
	}

	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err = fs.repo.Store(fileName, data)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			c.String(http.StatusConflict, "file already exists")
		} else {
			c.String(http.StatusInternalServerError, err.Error())
		}
		return
	}

	c.Status(http.StatusOK)
}

func (fs *FileStore) HandleDelete(c *gin.Context) {
	fileName := c.Query("fileName")
	if fileName == "" {
		c.String(http.StatusBadRequest, "fileName is required")
		return
	}

	err := fs.repo.Delete(fileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			c.String(http.StatusNotFound, "file not found")
		} else {
			c.String(http.StatusInternalServerError, err.Error())
		}
		return
	}

	c.Status(http.StatusOK)
}

func (fs *FileStore) HandleUpdate(c *gin.Context) {
	filename := c.Query("filename")
	if filename == "" {
		c.String(http.StatusBadRequest, "filename is required")
		return
	}

	data, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	err = fs.repo.Update(filename, data)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			c.String(http.StatusNotFound, "file not found")
		} else {
			c.String(http.StatusInternalServerError, err.Error())
		}
		return
	}

	c.Status(http.StatusOK)
}
