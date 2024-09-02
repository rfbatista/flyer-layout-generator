package layoutgenerator

import (
	"algvisual/internal/domain/entities"
	"algvisual/internal/infrastructure/config"
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

type GenerateImageRequestV2 struct {
	DesignFile string             `json:"design_file,omitempty"`
	Prancheta  entities.LayoutDTO `json:"prancheta,omitempty"`
}

type GenerateImageResultV2Elements struct {
	ElementID int32  `json:"element_id,omitempty"`
	ImageURL  string `json:"image_url,omitempty"`
}

type GenerateImageResultV2 struct {
	ImageURL string                          `json:"image_url,omitempty"`
	Elements []GenerateImageResultV2Elements `json:"elements,omitempty"`
}

func GenerateImageFromPranchetaV2(
	input GenerateImageRequestV2, log *zap.Logger, config config.AppConfig,
) (*GenerateImageResultV2, error) {
	jsonBody, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	// log.Info("generating image", zap.String("body", string(jsonBody)))
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
	var result GenerateImageResultV2
	json.NewDecoder(res.Body).Decode(&result)
	res2B, _ := json.Marshal(result)
	fmt.Println(string(res2B))
	return &result, nil
}
