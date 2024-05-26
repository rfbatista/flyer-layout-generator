package ports

import (
	entities2 "algvisual/internal/entities"
)

type ProcessFileResult struct {
	Photoshop entities2.DesignFile      `json:"photoshop,omitempty"`
	ImageUrl  string                    `json:"image_url,omitempty"`
	Elements  []entities2.DesignElement `json:"elements,omitempty"`
	Error     string                    `json:"error,omitempty"`
	Detail    string                    `json:"detail,omitempty"`
}

type ProcessFileInput struct {
	Filepath string `json:"filepath,omitempty"`
	ID       int32  `json:"id,omitempty"`
}

type PhotoshopProcessorServiceProcessFile func(input ProcessFileInput) (*ProcessFileResult, error)
