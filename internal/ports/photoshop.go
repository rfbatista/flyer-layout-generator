package ports

import "algvisual/internal/entities"

type ProcessFileResult struct {
	Photoshop entities.Photoshop          `json:"photoshop,omitempty"`
	Elements  []entities.PhotoshopElement `json:"elements,omitempty"`
	Error     string                      `json:"error,omitempty"`
}

type ProcessFileInput struct {
	Filepath string `json:"filepath,omitempty"`
}

type PhotoshopProcessorServiceProcessFile func(input ProcessFileInput) (*ProcessFileResult, error)
