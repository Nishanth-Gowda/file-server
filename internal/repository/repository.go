package repository

import (
	"errors"
	"log"
	"os"
	"path/filepath"
)

type Repository interface {
	Store(filename string, data []byte) error
	Update(filename string, data []byte) error
	Delete(filename string) error
}

type FileRepository struct {
	dataDir string
}

func NewFileRepository(dataDir string) (*FileRepository, error) {

	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return nil, err
	}
	return &FileRepository{
		dataDir: dataDir,
	}, nil
}

func (f *FileRepository) Store(fileName string, data []byte) error {
	filePath := filepath.Join(f.dataDir, fileName)

	//check if the file exists already
	_, err := os.Stat(filePath)
	if err != nil {
		return errors.New("file already exists")
	} else if !os.IsNotExist(err) {
		return err
	}

	if err := os.WriteFile(fileName, data, 0644); err != nil {
		return err
	}

	log.Println("File uploaded successfully")

	return nil

}

func (f *FileRepository) Update(fileName string, data []byte) error {
	filePath := filepath.Join(f.dataDir, fileName)

	// Check if the file exists
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return os.ErrNotExist
		}
		return err
	}

	if err := os.WriteFile(fileName, data, 0644); err != nil {
		return err
	}

	log.Println("File uploaded successfully")

	return nil

}

func (r *FileRepository) Delete(filename string) error {
	filePath := filepath.Join(r.dataDir, filename)

	// Check if the file exists
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// File doesn't exist
			return os.ErrNotExist
		}
		// Other error
		return err
	}

	if err := os.Remove(filePath); err != nil {
		return err
	}

	return nil
}

