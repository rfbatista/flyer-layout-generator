package infra

import (
	"fmt"
	"io"
	"os"
)

type Storage interface {
	Upload(file io.Reader, name string) (string, error)
}

func NewFileStorage(c *AppConfig) Storage {
	return FileStorage{
		dirpath: c.PhotoshopFilesPath,
	}
}

type FileStorage struct {
	dirpath string
}

func (f FileStorage) Upload(file io.Reader, name string) (string, error) {
	fpath := fmt.Sprintf("%s/%s", f.dirpath, name)
	dst, err := os.Create(
		fpath,
	) // dir is directory where you want to save file.
	if err != nil {
		return "", err
	}
	defer dst.Close()
	if _, err = io.Copy(dst, file); err != nil {
		return "", err
	}
	return fpath, nil
}
