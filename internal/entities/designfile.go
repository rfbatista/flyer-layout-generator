package entities

import (
	"strconv"
	"time"
)

type DesignFile struct {
	ID             int32     `json:"id,omitempty"`
	ProjectID      int32     `json:"project_id,omitempty"`
	Name           string    `json:"name,omitempty"`
	Filepath       string    `json:"filepath,omitempty"`
	LayoutID       int32     `json:"layout_id,omitempty"`
	FileExtension  string    `json:"file_extension,omitempty"`
	ImagePath      string    `json:"image_path,omitempty"`
	ImageURL       string    `json:"image_url,omitempty"`
	ImageExtension string    `json:"image_extension,omitempty"`
	Width          int32     `json:"width,omitempty"`
	Height         int32     `json:"height,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	IsProcessed    bool      `json:"is_processed,omitempty"`
}

func (d *DesignFile) SID() string {
	return strconv.FormatInt(int64(d.ID), 10)
}

func (d *DesignFile) SWidth() string {
	return strconv.FormatInt(int64(d.Width), 10)
}

func (d *DesignFile) SHeigth() string {
	return strconv.FormatInt(int64(d.Height), 10)
}
