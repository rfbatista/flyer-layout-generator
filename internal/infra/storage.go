package infra

import (
	"algvisual/internal/infra/config"
	"errors"
	"fmt"
	"image"
	"image/png"
	"io"
	"net/http"
	"os"
)

func NewFileStorage(c *config.AppConfig) FileStorage {
	return FileStorage{
		dirpath: c.PhotoshopFilesPath,
		cfg:     c,
	}
}

type FileStorage struct {
	dirpath string
	cfg     *config.AppConfig
}

func (f FileStorage) Upload(file io.Reader, name string) (string, error) {
	fpath := fmt.Sprintf("%s/%s.png", f.dirpath, name)
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

func (f FileStorage) LoadImageFromURL(URL string) (image.Image, error) {
	// Get the response bytes from the url
	response, err := http.Get(fmt.Sprintf("http://localhost:8000%s", URL))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New("received non 200 response code")
	}

	img, _, err := image.Decode(response.Body)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func (f FileStorage) SaveImage(name string, img image.Image) (string, error) {
	fullpath := fmt.Sprintf("%s/%s.png", f.cfg.ImagesFolderPath, name)
	dst, err := os.Create(fullpath)
	if err != nil {
		return "", err
	}
	defer dst.Close()
	fpath := fmt.Sprintf("%s/%s.png", f.dirpath, name)
	defer dst.Close()
	err = png.Encode(dst, img)
	if err != nil {
		return "", err
	}
	return fpath, nil
}
