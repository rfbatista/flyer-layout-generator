package usecase

import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type CreateZipForBatchInput struct{}

type CreateZipForBatchOutput struct{}

func CreateZipForBatchUseCase(
	ctx context.Context,
	req CreateZipForBatchInput,
) (*CreateZipForBatchOutput, error) {
	return &CreateZipForBatchOutput{}, nil
}

func downloadImage(url string) ([]byte, error) {
	resp, err := http.Get(url)
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
