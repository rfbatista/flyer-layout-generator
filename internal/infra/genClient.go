package infra

import (
	"algvisual/internal/entities"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
)

func NewImageGenerator(c *AppConfig, log *zap.Logger) (*ImageGeneratorClient, error) {
	return &ImageGeneratorClient{c: c, log: log}, nil
}

type GeneratorRequest struct {
	Template   entities.Template          `json:"template"`
	Photoshop  entities.DesignFile        `json:"photoshop"`
	Components []entities.DesignComponent `json:"components"`
	Elements   []entities.DesignElement   `json:"elements"`
}

type GeneratorResult struct {
	PhotoshopID int       `json:"photoshop_id,omitempty"`
	ImageURL    string    `json:"image_url,omitempty"`
	StartedAt   time.Time `json:"started_at,omitempty"`
	FinishedAt  time.Time `json:"finished_at,omitempty"`
	Logs        string    `json:"logs,omitempty"`
}

type ImageGeneratorClient struct {
	c   *AppConfig
	log *zap.Logger
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
	c.log.Info("generating image", zap.String("body", string(jsonBody)))
	bodyReader := bytes.NewReader(jsonBody)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*90)
	defer cancel()
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		c.c.AiServiceBaseURL+"/api/v1/generate/distortion",
		bodyReader,
	)
	if err != nil {
		c.log.Error("client: error making http request: %s\n", zap.Error(err))
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{
		Timeout: 90 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		c.log.Error("client: error doing http request: %s\n", zap.Error(err))
		os.Exit(1)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		buf := new(strings.Builder)
		io.Copy(buf, res.Body)
		c.log.Error(
			"error processing photoshop file",
			zap.Int("StatusCode", res.StatusCode),
			zap.Error(err),
		)
		err = fmt.Errorf("falha ao requisitar processamento do arquivo photoshop %s", buf.String())
		c.log.Error(err.Error())
		return nil, err
	}
	var result GeneratorResult
	json.NewDecoder(res.Body).Decode(&result)
	return &result, nil
}

type GenerateImageRequest struct {
	DesignFile string             `json:"design_file,omitempty"`
	Prancheta  entities.Prancheta `json:"prancheta,omitempty"`
}

type GenerateImageResult struct {
	ImageURL string `json:"image_url,omitempty"`
}

func GenerateImageFromPrancheta(
	input GenerateImageRequest, log *zap.Logger, config AppConfig,
) (*GenerateImageResult, error) {
	jsonBody, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	log.Info("generating image", zap.String("body", string(jsonBody)))
	bodyReader := bytes.NewReader(jsonBody)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*90)
	defer cancel()
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		config.AiServiceBaseURL+"/api/v1/prancheta/generate",
		bodyReader,
	)
	if err != nil {
		log.Error("client: error making http request: %s\n", zap.Error(err))
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{
		Timeout: 90 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		log.Error("client: error doing http request: %s\n", zap.Error(err))
		os.Exit(1)
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		buf := new(strings.Builder)
		io.Copy(buf, res.Body)
		log.Error(
			"error creating images",
			zap.Int("StatusCode", res.StatusCode),
			zap.Error(err),
		)
		err = fmt.Errorf("falha ao requisitar processamento do arquivo photoshop %s", buf.String())
		log.Error(err.Error())
		return nil, err
	}
	var result GenerateImageResult
	json.NewDecoder(res.Body).Decode(&result)
	return &result, nil
}
