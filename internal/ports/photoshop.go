package ports

import "algvisual/internal/entities"

type ProcessFileResult struct {
	Photoshop entities.DesignFile          `json:"photoshop,omitempty"`
	Elements  []entities.DesignElement `json:"elements,omitempty"`
	Error     string                      `json:"error,omitempty"`
	Detail    string
}

type ProcessFileInput struct {
	Filepath string `json:"filepath,omitempty"`
}

type PhotoshopProcessorServiceProcessFile func(input ProcessFileInput) (*ProcessFileResult, error)
