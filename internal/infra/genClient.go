package infra

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"algvisual/internal/entities"
)

func NewImageGenerator(c *AppConfig) (*ImageGeneratorClient, error) {
	return &ImageGeneratorClient{c: c}, nil
}

type GeneratorRequest struct {
	Template   entities.Template             `json:"template,omitempty"`
	Photoshop  entities.Photoshop            `json:"photoshop,omitempty"`
	Components []entities.PhotoshopComponent `json:"components,omitempty"`
	Elements   []entities.PhotoshopElement   `json:"elements,omitempty"`
}

type GeneratorResult struct {
	PhotoshopID int
	ImageURL    string
	StartedAt   time.Time
	FinishedAt  time.Time
	Logs        string
}

type ImageGeneratorClient struct {
	c *AppConfig
}

func (c ImageGeneratorClient) GenerateImageWithSlotStrategy(
	input GeneratorRequest,
) (*GeneratorResult, error) {
	jsonBody, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest(http.MethodPost, c.c.GeneratorClientURL, bodyReader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}
	defer res.Body.Close()
	var result GeneratorResult
	json.NewDecoder(res.Body).Decode(&result)
	return &result, nil
}

func (c ImageGeneratorClient) GenerateImageWithDistortionStrategy(
	input GeneratorRequest,
) (*GeneratorResult, error) {
	jsonBody, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest(
		http.MethodPost,
		c.c.GeneratorClientURL+"/api/v1/generate/distortion",
		bodyReader,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		os.Exit(1)
	}
	defer res.Body.Close()
	var result GeneratorResult
	json.NewDecoder(res.Body).Decode(&result)
	return &result, nil
}
