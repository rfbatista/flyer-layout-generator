package ports

import "io"

type Storage interface {
	Upload(file io.Reader, name string) (string, error)
}

type StorageUpload func(file io.Reader, name string) (string, error)
