package entities

import "time"

type DesignFile struct {
	ID             int32     `json:"id,omitempty"`
	Name           string    `json:"name,omitempty"`
	Filepath       string    `json:"filepath,omitempty"`
	FileExtension  string    `json:"file_extension,omitempty"`
	ImagePath      string    `json:"image_path,omitempty"`
	ImageExtension string    `json:"image_extension,omitempty"`
	Width          int32     `json:"width,omitempty"`
	Height         int32     `json:"height,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
}
