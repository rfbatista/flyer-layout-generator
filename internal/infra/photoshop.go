package infra

import (
	"algvisual/internal/ports"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
)

func NewPhotoshpProcessor(l *zap.Logger, c *AppConfig) (*PhotoshopProcessor, error) {
	return &PhotoshopProcessor{log: l, conf: c}, nil
}

type PhotoshopProcessor struct {
	log  *zap.Logger
	conf *AppConfig
}

func (p PhotoshopProcessor) ProcessFile(
	input ports.ProcessFileInput,
) (*ports.ProcessFileResult, error) {
	jsonBody, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	bodyReader := bytes.NewReader(jsonBody)
	url := fmt.Sprintf("%s%s", p.conf.AiServiceBaseURL, "/api/v1/photoshop")
	p.log.Info(fmt.Sprintf("calling photoshop file processor in %s", url))
	req, err := http.NewRequest(
		http.MethodPost,
		url,
		bodyReader,
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{
		Timeout: 90 * time.Second,
	}
	res, err := client.Do(req)
	p.log.Info("file processing finished")
	if err != nil {
		p.log.Error(
			"error processing photoshop file",
			zap.Error(err),
		)
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		buf := new(strings.Builder)
		io.Copy(buf, res.Body)
		p.log.Error(
			"error processing photoshop file",
			zap.Int("StatusCode", res.StatusCode),
			zap.Error(err),
		)
		err = fmt.Errorf("falha ao requisitar processamento do arquivo photoshop %s", buf.String())
		p.log.Error(err.Error())
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Failed to read response body: %s\n", err)
		return nil, err
	}
	var result ports.ProcessFileResult
	log.Info(fmt.Sprintf("response: %s", string(body)))
	err = json.Unmarshal(body, &result)
	if err != nil {
		p.log.Warn("error ao desempacotar resultados do arquivo processado", zap.Error(err))
		return nil, err
	}
	return &result, nil
}
