package infra

import (
	"algvisual/internal/infra/config"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewFileLogger() {}

func NewLogger(cfg *config.AppConfig) *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	if cfg.IsProd() {
		config = zap.NewProductionConfig()
		fullpath := filepath.Join(cfg.LogPath, "proxy.log")
		config.Level = zap.NewAtomicLevel()
		config.OutputPaths = []string{
			fullpath,
		}
	}
	logger, _ := config.Build()
	return logger
}

func NewTestLogger() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, _ := config.Build()
	return logger
}
