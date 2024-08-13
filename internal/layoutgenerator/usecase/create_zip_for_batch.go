package usecase

import (
	"algvisual/internal/layoutgenerator/repository"
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type CreateZipForBatchInput struct {
	RequestID int64
}

type CreateZipForBatchOutput struct {
	Data []byte
}

func CreateZipForBatchUseCase(
	ctx context.Context,
	req CreateZipForBatchInput,
	repo repository.LayoutRepository,
) (*CreateZipForBatchOutput, error) {
	layouts, err := repo.GetLayoutByRequestID(ctx, req.RequestID)
	if err != nil {
		return nil, err
	}
	var urls []string
	for _, l := range layouts {
		urls = append(urls, l.ImageURL)
	}
	images, err := zipImages("teste", urls)
	if err != nil {
		return nil, err
	}
	return &CreateZipForBatchOutput{
		Data: images,
	}, nil
}

func downloadImage(url string) ([]byte, error) {
	resp, err := http.Get(fmt.Sprintf("http://localhost:8000%s", url))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func zipImages(name string, urls []string) ([]byte, error) {
	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)

	for _, url := range urls {
		data, err := downloadImage(url)
		if err != nil {
			return nil, err
		}

		fileName := fmt.Sprintf("%s::%s", name, url[strings.LastIndex(url, "/")+1:])
		fileWriter, err := zipWriter.Create(fileName)
		if err != nil {
			return nil, err
		}

		_, err = fileWriter.Write(data)
		if err != nil {
			return nil, err
		}
	}

	err := zipWriter.Close()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
