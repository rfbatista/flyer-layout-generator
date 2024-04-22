package infra

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"go.uber.org/zap"
)

func NewPhotoshpProcessor(l *zap.Logger, c *AppConfig) (*PhotoshopProcessor, error) {
	return &PhotoshopProcessor{log: l, conf: c}, nil
}

type PhotoshopProcessor struct {
	log  *zap.Logger
	conf *AppConfig
}

type photoshop struct {
	Filepath string `json:"filepath,omitempty"`
	Width    int    `json:"width,omitempty"`
	Height   int    `json:"height,omitempty"`
}

type photoshopElement struct {
	Xi          int    `json:"xi,omitempty"`
	Xii         int    `json:"xii,omitempty"`
	Yi          int    `json:"yi,omitempty"`
	Yii         int    `json:"yii,omitempty"`
	LayerID     string `json:"layer_id,omitempty"`
	Width       int    `json:"width,omitempty"`
	Height      int    `json:"height,omitempty"`
	Kind        string `json:"kind,omitempty"`
	Name        string `json:"name,omitempty"`
	IsGroup     bool   `json:"is_group,omitempty"`
	GroupId     int    `json:"group_id,omitempty"`
	Level       int    `json:"level,omitempty"`
	PhotoshopId int    `json:"photoshop_id,omitempty"`
	Image       string `json:"image,omitempty"`
	Text        string `json:"text,omitempty"`
}

type processFileResult struct {
	Photoshop photoshop          `json:"photoshop,omitempty"`
	Elements  []photoshopElement `json:"elements,omitempty"`
	Error     string             `json:"error,omitempty"`
}

type ProcessFileInput struct {
	Filepath string `json:"filepath,omitempty"`
}

func (p PhotoshopProcessor) ProcessFile(input ProcessFileInput) (*processFileResult, error) {
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
		Timeout: 30 * time.Second,
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
	statusOK := res.StatusCode >= 200 && res.StatusCode < 300
	if !statusOK {
		buf := new(strings.Builder)
		io.Copy(buf, res.Body)
		err = fmt.Errorf("falha ao requisitar processamento do arquivo photoshop %s", buf.String())
		p.log.Error(err.Error())
		// You may read / inspect response body
		return nil, err
	}
	var result processFileResult
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		p.log.Warn("error ao desempacotar resultados do arquivo processado", zap.Error(err))
		return nil, err
	}
	return &result, nil
}
